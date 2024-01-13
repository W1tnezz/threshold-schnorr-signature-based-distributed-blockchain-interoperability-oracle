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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("marshal R_i error : %v", err)
	}

	log.Infof("Start send kafka message R")
	v.sendR(RiBytes)

	// 此时需要获取到其他人的R,此时需要等待其他人广播完成，获取完全足够的R
	timeout := time.After(Timeout)
	count, err := v.oracleContract.CountEnrollNodes(nil)
loop:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop
		default:
			if count.Int64() == int64(len(v.RAll)) {
				break loop
			}
			time.Sleep(50 * time.Millisecond)
		}
	}

	R := v.suite.G1().Point().Null()
	for key := range v.RAll {
		R = v.suite.G1().Point().Add(R, v.RAll[key])
	}

	lam, Y := v.suite.G1().Scalar().Pick(v.suite.RandomStream()), v.suite.G1().Point().Pick(v.suite.RandomStream())
	m := message
	RByte, err := R.MarshalBinary()
	if err != nil {
		log.Errorf("marshal R error : %w", err)
	}
	YBytes, err := Y.MarshalBinary()
	if err != nil {
		log.Errorf("marshal Y error : %w", err)
	}

	m = append(m, RByte...)
	m = append(m, YBytes...)

	hash := sha256.New()
	hash.Write(m)
	c := v.suite.G1().Scalar().SetBytes(hash.Sum(nil))

	signature := make([][]byte, 2)
	z := v.suite.G1().Scalar().Add(ri, v.suite.G1().Scalar().Mul(lam, v.suite.G1().Scalar().Mul(c, v.privateKey)))

	signature[0], err = z.MarshalBinary()
	if err != nil {
		log.Errorf("marshal z error : %w", err)
	}
	signature[1] = RByte

	return signature, nil

}

// func (v *Validator) SignForBls(message []byte, enrollNodes []int64) ([][]byte, error) {
// 	hash := sha256.New()
// 	hash.Write(message)

// 	messageHash := hash.Sum(nil)

// 	_hash := v.suite.G1().Point().Mul(v.suite.G1().Scalar().SetBytes(messageHash), nil)

// 	signature := make([][]byte, 2)
// 	node, _ := v.oracleContract.FindOracleNodeByAddress(nil, v.account)

// 	for i := int64(0); i < v.reputation; i++ {
// 		sI := v.suite.G1().Point().Mul(v.privateKey[i], _hash)
// 		pubkey := node.BlsPubKeys[i]
// 		fmt.Println("141", pubkey)

// 		PKbytes := make([]byte, 0)

// 		for _, z := range [4]int{1, 0, 3, 2} {

// 			sub := 32 - len(pubkey[z].Bytes())

// 			bigByte := make([]byte, sub)

// 			// for i := 0; i < sub; i++ {
// 			// 	bigByte = append(bigByte, 0)
// 			// }

// 			bigByte = append(bigByte, pubkey[z].Bytes()...)
// 			PKbytes = append(PKbytes, bigByte...)

// 		}
// 		pk := v.suite.G2().Point()
// 		err := pk.UnmarshalBinary(PKbytes)
// 		if err != nil {
// 			fmt.Println("161 translate pk", err)
// 		}
// 		fmt.Println("139", v.suite.Pair(_hash, pk).Equal(v.suite.Pair(sI, v.suite.G2().Point().Base())))
// 		siBytes, _ := sI.MarshalBinary()
// 		for _, b := range siBytes {
// 			signature[0] = append(signature[0], b)
// 		}
// 	}
// 	return signature, nil

// }

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
					log.Errorf("R transform to Point: %v", err)
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

func (v *Validator) ValidateTransaction(ctx context.Context, hash common.Hash, size int64, minRank int64) (*ValidateResult, error) {
	log.Info("请求 receipt")
	receipt, err := v.ethClient.TransactionReceipt(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil {
		return nil, fmt.Errorf("transaction receipt: %w", err)
	}
	log.Info("请求 blocknumber")
	blockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	valid := true
	if found {
		confirmed := blockNumber - receipt.BlockNumber.Uint64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid, ValidateRequest_transaction)
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
	}, nil
}

func (v *Validator) ValidateBlock(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
	block, err := v.ethClient.BlockByHash(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil && found {
		return nil, fmt.Errorf("block: %w", err)
	}

	latestBlockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	var blockNumber *big.Int
	valid := false
	if found {
		blockNumber = block.Number()
		confirmed := latestBlockNumber - block.NumberU64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid, ValidateRequest_block)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}

	// distKey, err := v.dkg.DistKeyShare()
	if err != nil {
		return nil, fmt.Errorf("dist key share: %w", err)
	}

	sig, err := v.Sign(message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %w", err)
	}

	return &ValidateResult{
		hash,
		valid,
		blockNumber,
		sig[0],
		sig[1],
	}, nil
}
