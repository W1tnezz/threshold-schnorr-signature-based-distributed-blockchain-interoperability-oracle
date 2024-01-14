package node

import (
	"context"

	"crypto/ecdsa"
	"crypto/sha256"

	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
)

const CONFIRMATIONS uint64 = 0

type ValidateResult struct {
	hash        common.Hash
	valid       bool
	blockNumber *big.Int
	signature   []byte
	R           []byte
	message     []byte
}

type Validator struct {
	sync.RWMutex
	suite           pairing.Suite
	oracleContract  *OracleContractWrapper
	ecdsaPrivateKey *ecdsa.PrivateKey
	ethClient       *ethclient.Client
	RAll            map[common.Address]kyber.Point
	account         common.Address
	kafkaWriter     *kafka.Writer
	kafkaReader     *kafka.Reader
	privateKey      kyber.Scalar
	enrolled        bool
}

func NewValidator(
	suite pairing.Suite,
	oracleContract *OracleContractWrapper,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	RAll map[common.Address]kyber.Point,
	account common.Address,
	kafkaWriter *kafka.Writer,
	kafkaReader *kafka.Reader,
	privateKey kyber.Scalar,
	enrolled bool,
) *Validator {
	return &Validator{
		suite:           suite,
		ecdsaPrivateKey: ecdsaPrivateKey,
		oracleContract:  oracleContract,
		ethClient:       ethClient,
		RAll:            RAll,
		account:         account,
		kafkaWriter:     kafkaWriter,
		kafkaReader:     kafkaReader,
		privateKey:      privateKey,
		enrolled:        enrolled,
	}
}

func (v *Validator) Sign(message []byte) ([][]byte, error) {
	// 先产生自己的R，然后在等待一段时间，随后广播, 构造R序列
	ri := v.suite.G1().Scalar().Pick(random.New())
	Ri := v.suite.G1().Point().Mul(ri, nil)

	RiBytes, err := Ri.MarshalBinary()

	if err != nil {
		log.Println("marshal R_i error : ", err)
	}

	log.Println("Start send kafka message R")
	v.sendR(RiBytes)

	// 此时需要获取到其他人的R,此时需要等待其他人广播完成，获取完全足够的R
	timeout := time.After(Timeout)
	validators, err := v.oracleContract.DKG.GetValidators(nil)
loop:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop
		default:
			if len(validators) == len(v.RAll) {
				break loop
			}
			time.Sleep(50 * time.Millisecond)
		}
	}

	R := v.suite.G1().Point().Null()
	for key := range v.RAll {
		R = v.suite.G1().Point().Add(R, v.RAll[key])
	}
	lamBig, err := v.oracleContract.Registry.GetLambda(nil, v.account)
	if err != nil {
		log.Println("get lam err : ", err)
	}

	YBig, err := v.oracleContract.DKG.GetPubKey(nil)
	if err != nil {
		log.Println("get Y err : ", err)
	}

	lam := v.suite.G1().Scalar().SetBytes(lamBig.Bytes())

	m := message
	RByte, err := R.MarshalBinary()
	if err != nil {
		log.Println("marshal R error : ", err)
	}

	m = append(m, RByte...)
	m = append(m, YBig[0].Bytes()...)
	m = append(m, YBig[1].Bytes()...)

	hash := sha256.New()
	hash.Write(m)
	c := v.suite.G1().Scalar().SetBytes(hash.Sum(nil))

	signature := make([][]byte, 2)
	z := v.suite.G1().Scalar().Add(ri, v.suite.G1().Scalar().Mul(lam, v.suite.G1().Scalar().Mul(c, v.privateKey)))

	signature[0], err = z.MarshalBinary()
	if err != nil {
		log.Println("marshal z error : ", err)
	}
	signature[1] = RiBytes

	// pk := v.suite.G1().Point().Mul(lam, v.suite.G1().Point().Mul(v.privateKey, nil))
	// pkByte, _ := pk.MarshalBinary()
	// signature[1] = pkByte
	return signature, nil

}

func (v *Validator) ListenAndProcess(o *OracleNode) error {

	for {
		m, err := v.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		// 处理kafka消息
		if v.enrolled {
			go func() {
				RPoint := v.suite.G1().Point()
				err := RPoint.UnmarshalBinary(m.Value)
				if err != nil {
					log.Println("R transform to Point: ", err)
				}
				v.RAll[common.Address(m.Key)] = RPoint
			}()
		}
	}
	return nil
}

func (v *Validator) sendR(R []byte) {
	messages := []kafka.Message{
		{
			Key:   []byte(v.account.String()),
			Value: R,
		},
	}
	var err error
	const retries = 3
	// 重试3次

	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = v.kafkaWriter.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
		break
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
	log.Println("请求 receipt")
	receipt, err := v.ethClient.TransactionReceipt(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil {
		return nil, fmt.Errorf("transaction receipt: %w", err)
	}
	log.Println("请求 blocknumber")
	blockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	valid := true
	if found {
		confirmed := blockNumber - receipt.BlockNumber.Uint64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("dist key share: %w", err)
	}

	// 以下是进行签名，

	sig, err := v.Sign(message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %w", err)
	}

	return &ValidateResult{
		hash,
		valid,
		big.NewInt(int64(blockNumber)),
		sig[0],
		sig[1],
		message,
	}, nil
}

// func (v *Validator) ValidateBlock(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
// 	block, err := v.ethClient.BlockByHash(ctx, hash)
// 	found := !errors.Is(err, ethereum.NotFound)
// 	if err != nil && found {
// 		return nil, fmt.Errorf("block: %w", err)
// 	}

// 	latestBlockNumber, err := v.ethClient.BlockNumber(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("blocknumber: %w", err)
// 	}

// 	var blockNumber *big.Int
// 	valid := false
// 	if found {
// 		blockNumber = block.Number()
// 		confirmed := latestBlockNumber - block.NumberU64()
// 		valid = confirmed >= CONFIRMATIONS
// 	}

// 	message, err := encodeValidateResult(hash, valid, ValidateRequest_block)
// 	if err != nil {
// 		return nil, fmt.Errorf("encode result: %w", err)
// 	}

// 	// distKey, err := v.dkg.DistKeyShare()
// 	if err != nil {
// 		return nil, fmt.Errorf("dist key share: %w", err)
// 	}

// 	sig, err := v.Sign(message)
// 	if err != nil {
// 		return nil, fmt.Errorf("tbls sign: %w", err)
// 	}

// 	return &ValidateResult{
// 		hash,
// 		valid,
// 		blockNumber,
// 		sig[0],
// 		sig[1],
// 	}, nil
// }
