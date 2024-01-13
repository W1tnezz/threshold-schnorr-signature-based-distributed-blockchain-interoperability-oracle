package node

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Aggregator struct {
	suite     pairing.Suite
	ethClient *ethclient.Client

	oracleContract  *OracleContractWrapper
	account         common.Address
	ecdsaPrivateKey *ecdsa.PrivateKey
	chainId         *big.Int
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	oracleContract *OracleContractWrapper,
	account common.Address,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	chainId *big.Int,

) *Aggregator {
	return &Aggregator{
		suite:           suite,
		ethClient:       ethClient,
		oracleContract:  oracleContract,
		account:         account,
		ecdsaPrivateKey: ecdsaPrivateKey,
		chainId:         chainId,
	}
}

func (a *Aggregator) WatchAndHandleValidationRequestsLog(ctx context.Context, o *OracleNode) error {
	sink := make(chan *OracleContractValidationRequest)
	defer close(sink)

	sub, err := a.oracleContract.WatchValidationRequest(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received ValidationRequest event with hash %s", common.Hash(event.Hash))
			isAggregator, err := a.oracleContract.IsAggregator(nil, a.account)
			o.isAggregator = isAggregator
			if err != nil {
				log.Errorf("Is aggregator: %v", err)
				continue
			}

			if !isAggregator && event.NeedEnroll {
				// 报名函数
				// node, err := a.registryContract.FindOracleNodeByAddress(nil, a.account)
				// time.Sleep(time.Duration(node.Index.Int64()) * time.Second)

				err = a.Enroll()
				if err != nil {
					log.Errorf("Node Enroll log: %v", err)
				} else {
					o.validator.enrolled = true
					log.Infof("Enroll success")
				}
				continue
			}


			if err := a.HandleValidationRequest(ctx, event); err != nil {
				log.Errorf("Handle ValidationRequest log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// 报名函数
func (a *Aggregator) Enroll() error {
	
	auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
	_, err = a.oracleContract.DKG.Enroll(auth)
	if err != nil {
		return fmt.Errorf("enroll iop node: %w", err)
	}
	return nil
}

func (a *Aggregator) WatchAndHandleDKGLog(ctx context.Context, event *OracleContractValidationRequest) error {
	sink := make(chan *DKGDistKey)
	defer close(sink)

	sub, err := a.oracleContract.DKG.WatchDistKey(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:

		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (a *Aggregator) HandleValidationRequest(ctx context.Context, event *OracleContractValidationRequest) error {

	

	result, MulSig, MulR, _hash, err := a.AggregateValidationResults(ctx, event.Hash) // schnorr
	// result, MulSig, _hash, MulY, nodes, pkSet, err := a.AggregateValidationResults(ctx, event.Hash, typ)

	if err != nil {
		return fmt.Errorf("aggregate validation results: %w", err)
	}
	if err != nil {
		return fmt.Errorf("signature to big int: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	sig, err := ScalarToBig(MulSig) // schnorr
	// sig, err := G1PointToBig(MulSig) // bls
	fmt.Println(sig)

	if err != nil {
		return fmt.Errorf("signature tranform to big int: %w", err)
	}
	if err != nil {
		return fmt.Errorf("public key tranform to big int: %w", err)
	}
	R, err := G1PointToBig(MulR) // schnorr

	if err != nil {
		return fmt.Errorf("multi R tranform to big int: %w", err)
	}
	hash, err := ScalarToBig(_hash) //schnorr
	if err != nil {
		return fmt.Errorf("hash tranform to big int: %w", err)
	}

	_, err = a.oracleContract.OracleContract.Submit(auth, result, event.Hash, sig, R[0], R[1], hash)

	if err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}

	resultStr := "valid"
	if !result {
		resultStr = "invalid"
	}
	log.Infof("Submitted validation result (%s) for hash %s of type %s", resultStr, common.Hash(event.Hash))

	return nil
}

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash) (bool, kyber.Scalar, kyber.Point, kyber.Scalar, error) { // schnorr

	Signatures := make([]kyber.Scalar, 0)
	Rs := make([]kyber.Point, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	enrollNodes, err := a.oracleContract.GetValidators(nil)
	if err != nil {
		log.Error("get enrollNodes %w", err)
	}
	var m []byte
	for _, enrollNode := range enrollNodes {

		node, _ := a.oracleContract.Registry.GetNodeByAddress(nil, enrollNode)
		conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
		if err != nil {
			log.Errorf("Find connection by address: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			client := NewOracleNodeClient(conn)
			ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)

			result, err := client.Validate(ctxTimeout, &ValidateRequest{
				Hash: txHash[:],
			})

			cancel()
			if err != nil {
				log.Errorf("Validate %s: %v", err)
				return
			}

			mutex.Lock()
			if result.Valid {
				m = result.Message
				z := a.suite.G1().Scalar().SetBytes(result.Signature)
				R := a.suite.G1().Point().Null()
				R.UnmarshalBinary(result.R)

				Signatures = append(Signatures, z) //获取到所有的签名
				Rs = append(Rs, R)
			}
		}()
	}
	wg.Wait()

	R := a.suite.G1().Point().Null()
	MulSig := a.suite.G1().Scalar().Zero()
	for index, _ := range Rs {
		R = a.suite.G1().Point().Add(R, Rs[index])
		MulSig = a.suite.G1().Scalar().Add(MulSig, Signatures[index])
	}

	YBig, err := a.oracleContract.DKG.GetPubKey(nil)
	if err != nil {
		log.Errorf("get Y err : %w", err)
	}

	RByte, err := R.MarshalBinary()
	if err != nil {
		log.Errorf("marshal R error : %w", err)
	}

	m = append(m, RByte...)
	m = append(m, YBig[0].Bytes()...)
	m = append(m, YBig[1].Bytes()...)

	hash := sha256.New()
	hash.Write(m)
	c := a.suite.G1().Scalar().SetBytes(hash.Sum(nil))
	return true, MulSig, R, c, nil

}
