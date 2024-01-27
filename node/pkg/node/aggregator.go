package node

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"google.golang.org/grpc"

	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	DKGSuccess      bool
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	oracleContract *OracleContractWrapper,
	account common.Address,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	chainId *big.Int,
	DKGSuccess bool,
) *Aggregator {
	return &Aggregator{
		suite:           suite,
		ethClient:       ethClient,
		oracleContract:  oracleContract,
		account:         account,
		ecdsaPrivateKey: ecdsaPrivateKey,
		chainId:         chainId,
		DKGSuccess:      DKGSuccess,
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
			log.Println("Received ValidationRequest event with hash ", common.Hash(event.Hash))
			isAggregator, err := a.oracleContract.IsAggregator(nil, a.account)
			o.isAggregator = isAggregator
			if err != nil {
				log.Println("Is aggregator: ", err)
				continue
			}

			if !isAggregator && event.NeedEnroll {
				// 报名函数
				// node, err := a.registryContract.FindOracleNodeByAddress(nil, a.account)
				// time.Sleep(time.Duration(node.Index.Int64()) * time.Second)

				err = a.Enroll()
				if err != nil {
					log.Println("Node Enroll log: ", err)
				} else {
					o.validator.enrolled = true
					log.Println("Enroll success")
				}
				continue
			}

			for !a.DKGSuccess {

			}

			if err := a.HandleValidationRequest(ctx, event); err != nil {
				log.Println("Handle ValidationRequest log: ", err)
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

func (a *Aggregator) WatchAndHandleDKGLog(ctx context.Context) error {
	sink := make(chan *DKGDistKey)
	defer close(sink)

	sub, err := a.oracleContract.DKG.WatchDistKey(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			a.DKGSuccess = true
			_ = event
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (a *Aggregator) HandleValidationRequest(ctx context.Context, event *OracleContractValidationRequest) error {

	// result, MulSig, MulR, _hash, YBig, err := a.AggregateValidationResults(ctx, event.Hash) // schnorr
	result, MulSig, MulR, message, err := a.AggregateValidationResults(ctx, event.Hash) // schnorr

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
	// hash, err := ScalarToBig(_hash) //schnorr
	// _ = hash
	if err != nil {
		return fmt.Errorf("hash tranform to big int: %w", err)
	}

	_, err = a.oracleContract.OracleContract.Submit(auth, result, event.Hash, message, sig, R[0], R[1])

	if err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}

	resultStr := "valid"
	if !result {
		resultStr = "invalid"
	}
	log.Println("Submitted validation result () for hash  of type ", resultStr, common.Hash(event.Hash))

	return nil
}

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash) (bool, kyber.Scalar, kyber.Point, []byte, error) { // schnorr

	Signatures := make([]kyber.Scalar, 0)
	Rs := make([]kyber.Point, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	enrollNodes, err := a.oracleContract.GetValidators(nil)
	if err != nil {
		log.Println("get enrollNodes ", err)
	}
	var message []byte
	for _, enrollNode := range enrollNodes {

		node, _ := a.oracleContract.Registry.GetNodeByAddress(nil, enrollNode)
		conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
		if err != nil {
			log.Println("Find connection by address: ", err)
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
				log.Println("Validate err :", err)
				return
			}

			mutex.Lock()
			if result.Valid {
				message = result.Message
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

	return true, MulSig, R, message, nil

}
