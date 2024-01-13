package node

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sync"
	"time"

	"go.dedis.ch/kyber/v3/util/random"
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
			typ := ValidateRequest_Type(event.Typ)
			log.Infof("Received ValidationRequest event for %s type with hash %s", typ, common.Hash(event.Hash))
			isAggregator, err := a.oracleContract.IsAggregator(nil, a.account)
			o.isAggregator = isAggregator
			if err != nil {
				log.Errorf("Is aggregator: %v", err)
				continue
			}

			if !isAggregator && event.index%6 == 0 {
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

			if err := a.HandleValidationRequest(ctx, event, typ); err != nil {
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
	isEnroll, err := a.oracleContract.OracleNodeIsEnroll(nil, a.account)
	if err != nil {
		return fmt.Errorf("is enrolled: %w", err)
	}
	if !isEnroll {
		auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
		_, err = a.oracleContract.EnrollOracleNode(auth)
		if err != nil {
			return fmt.Errorf("enroll iop node: %w", err)
		}
	}
	return nil
}

func (a *Aggregator) HandleValidationRequest(ctx context.Context, event *OracleContractValidationRequest, typ ValidateRequest_Type) error {
	result, MulSig, MulR, _hash, MulY, nodes, pkSet, err := a.AggregateValidationResults(ctx, event.Hash, typ) // schnorr
	// result, MulSig, _hash, MulY, nodes, pkSet, err := a.AggregateValidationResults(ctx, event.Hash, typ)

	_ = MulR
	pk, err := G1PointToBig(MulY)
	_, _ = pk, pkSet

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
	// hash, err := G1PointToBig(_hash)
	if err != nil {
		return fmt.Errorf("hash tranform to big int: %w", err)
	}
	switch typ {
	case ValidateRequest_block:
		// _, err = a.oracleContract.SubmitBlockValidationResult(auth, result, event.Hash, big.NewInt(0), hash[0], hash[1], big.NewInt(0), nodes)
	case ValidateRequest_transaction:
		_, err = a.oracleContract.SubmitTransactionValidationResult(auth, result, event.Hash, sig, R[0], R[1], hash, nodes)
		// _, err = a.oracleContract.SubmitValidationResult(auth, 0, result, event.Hash, pk, sig, hash, nodes)
	default:
		return fmt.Errorf("unknown validation request type %s")
	}

	if err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}

	resultStr := "valid"
	if !result {
		resultStr = "invalid"
	}
	log.Infof("Submitted validation result (%s) for hash %s of type %s", resultStr, common.Hash(event.Hash), typ)

	return nil
}

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash, typ ValidateRequest_Type) (bool, kyber.Scalar, kyber.Point, kyber.Scalar, kyber.Point, []common.Address, [][2]*big.Int, error) { // schnorr

	Signatures := make([][]kyber.Scalar, 0)
	Rs := make([][]kyber.Point, 0)
	// Signatures := make([][]kyber.Point, 0)
	PK := make([][][2]*big.Int, 0)
	nodes := make([]common.Address, 0)
	totalRank := int64(0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	enrollNodes := make([]common.Address, 0)
	tmpScalar := a.suite.G1().Scalar().Pick(random.New())
	scalarSize := tmpScalar.MarshalSize()
	PointSize := a.suite.G1().Point().Mul(tmpScalar, nil).MarshalSize()

	for _, enrollNode := range enrollNodes {

		node, _ := a.oracleContract.FindOracleNodeByAddress(nil, enrollNode)
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
				Type: typ,
				Hash: txHash[:],
			})

			cancel()
			if err != nil {
				log.Errorf("Validate %s: %v", typ, err)
				return
			}

			mutex.Lock()
			if result.Valid {

				sI, RI := a.HandleResultForSchnorr(result, scalarSize, PointSize)
				// sI := a.HandleResultForBls(result, PointSize)

				Signatures = append(Signatures, sI) //获取到所有的签名
				Rs = append(Rs, RI)
			}
		}()
	}
	wg.Wait()

	return a.AggregateSignatureForSchnorr(txHash, typ, Signatures, Rs, PK, nodes, totalRank)
	// return a.AggregateSignatureForBLS(txHash, typ, Signatures, PK, nodes, totalRank)

}

func (a *Aggregator) HandleResultForSchnorr(result *ValidateResponse, scalarSize int, PointSize int) ([]kyber.Scalar, []kyber.Point) {
	sI := make([]kyber.Scalar, 0)
	RI := make([]kyber.Point, 0)

	for i := int64(0); i < result.Reputation; i++ {
		sSlice := result.Signature[i*int64(scalarSize) : (i+1)*int64(scalarSize)]

		sI = append(sI, a.suite.G1().Scalar().SetBytes(sSlice))

		RSliceBytes := result.R[i*int64(PointSize) : (i+1)*int64(PointSize)]
		RSlice := a.suite.G1().Point().Base()
		err := RSlice.UnmarshalBinary(RSliceBytes)
		if err != nil {
			fmt.Println("UnmarshalBinary R ,", err)
		}
		RI = append(RI, RSlice)
	}
	return sI, RI
}

func (a *Aggregator) HandleResultForBls(result *ValidateResponse, PointSize int) []kyber.Point {
	sI := make([]kyber.Point, 0)

	for i := int64(0); i < result.Reputation; i++ {
		sSliceByte := result.Signature[i*int64(PointSize) : (i+1)*int64(PointSize)]

		sSlice := a.suite.G1().Point()
		err := sSlice.UnmarshalBinary(sSliceByte)
		if err != nil {
			fmt.Println("UnmarshalBinary Si for Bls ,", err)
		}
		sI = append(sI, sSlice)
	}
	return sI
}

func (a *Aggregator) AggregateSignatureForSchnorr(txHash common.Hash, typ ValidateRequest_Type, Signatures [][]kyber.Scalar, Rs [][]kyber.Point, PK [][][2]*big.Int, nodes []common.Address, totalRank int64) (bool, kyber.Scalar, kyber.Point, kyber.Scalar, kyber.Point, []common.Address, [][2]*big.Int, error) {
	index := 64
	S := make([]byte, (totalRank+1)*64)

	pkSet := make([][2]*big.Int, 0)

	zeroPointBytes := make([]byte, 64)
	pointS := a.suite.G1().Point()
	pointS.UnmarshalBinary(zeroPointBytes)
	for i := 0; i < len(a.enrollNodes); i++ {
		for j := 0; j < len(PK[i]); j++ {
			// 构造Point累加形式的S
			PKbytes := make([]byte, 0)
			for k := 0; k < 32-len(PK[i][j][0].Bytes()); k++ {
				PKbytes = append(PKbytes, 0)
			}
			PKbytes = append(PKbytes, PK[i][j][0].Bytes()...)
			for k := 0; k < 32-len(PK[i][j][1].Bytes()); k++ {
				PKbytes = append(PKbytes, 0)
			}
			PKbytes = append(PKbytes, PK[i][j][1].Bytes()...)

			PKpoint := a.suite.G1().Point()
			PKpoint.UnmarshalBinary(PKbytes)
			pointS = a.suite.G1().Point().Add(pointS, PKpoint)

			pkSet = append(pkSet, PK[i][j])

			for k := 0; k < 2; k++ {
				tmp := PK[i][j][k].Bytes()
				for _, byteTmp := range tmp {
					S[index] = byteTmp
					index++
				}
			}
		}

	}

	MulSignature := a.suite.G1().Scalar().Zero()
	MulR := a.suite.G1().Point().Null()
	MulY := a.suite.G1().Point().Null()

	R := a.suite.G1().Point().Null()

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(PK[i]); j++ {

			tmpX := PK[i][j][0]
			tmpY := PK[i][j][1]
			tmpXByte := tmpX.Bytes()
			XByte := make([]byte, 32)

			for k := 31; k >= 0; k-- {
				if len(tmpXByte)-(len(XByte)-k) >= 0 {
					XByte[k] = tmpXByte[len(tmpXByte)-(len(XByte)-k)]
				} else {
					XByte[k] = 0
				}

			}
			for k := 0; k < 32; k++ {
				S[k] = XByte[k]
			}
			tmpYByte := tmpY.Bytes()

			YByte := make([]byte, 32)
			for k := 31; k >= 0; k-- {
				if len(tmpYByte)-(len(YByte)-k) >= 0 {
					YByte[k] = tmpYByte[len(tmpYByte)-(len(YByte)-k)]
				} else {
					YByte[k] = 0
				}

			}
			for k := 0; k < 32; k++ {
				S[k+32] = YByte[k]
			}
			pkbytes := S[0:64]
			pk := a.suite.G1().Point()
			err := pk.UnmarshalBinary(pkbytes)

			if err != nil {
				fmt.Println("translate pk ", err)
			}

			// 累加S
			summaryS := a.suite.G1().Point().Add(pointS, pk)
			summarySBytes, err := summaryS.MarshalBinary()
			if err != nil {
				log.Errorf("Marshal point err: %v", err)
			}

			hash1 := sha256.New()

			// hash1.Write(S)
			hash1.Write(summarySBytes)
			aI := hash1.Sum(nil)

			aScalar := a.suite.G1().Scalar().SetBytes(aI)
			MulSignature.Add(MulSignature, a.suite.G1().Scalar().Mul(aScalar, Signatures[i][j]))
			MulY.Add(MulY, a.suite.G1().Point().Mul(aScalar, pk))
			MulR.Add(MulR, a.suite.G1().Point().Mul(aScalar, Rs[i][j]))
			R.Add(R, Rs[i][j])
		}
	}

	message, _ := encodeValidateResult(txHash, true, typ)

	m := make([][]byte, 2)
	m[0] = message
	m[1], _ = R.MarshalBinary()
	hash := sha256.New()
	e := hash.Sum(bytes.Join(m, []byte("")))
	_hash := a.suite.G1().Scalar().SetBytes(e)

	left := a.suite.G1().Point().Mul(MulSignature, nil)
	right := MulR.Clone()

	right.Add(right, a.suite.G1().Point().Mul(_hash, MulY))
	fmt.Println("435", right.Equal(left))
	a.enrollNodes = []int64{}

	return true, MulSignature, MulR, _hash, MulY, nodes, pkSet, nil

}

func (a *Aggregator) AggregateSignatureForBLS(txHash common.Hash, typ ValidateRequest_Type, Signatures [][]kyber.Point, PK [][][4]*big.Int, nodes []common.Address, totalRank int64) (bool, kyber.Point, kyber.Point, kyber.Point, []common.Address, [][2]*big.Int, error) {
	pkSet := make([][4]*big.Int, 0)

	// PointBig, err := a.oracleContract.GetNodeBLSPublicKeysSum(nil)
	PointBig := new([4]*big.Int)

	// if err != nil {
	// 	fmt.Println("GetNodeBLSPublicKeysSum : ", err, PointBig)
	// }

	PointByte := make([]byte, 0)

	for _, z := range [4]int{1, 0, 3, 2} {
		bigByte := PointBig[z].Bytes()
		sub := 32 - len(bigByte)

		bigByte = make([]byte, 0)
		for i := 0; i < sub; i++ {
			bigByte = append(bigByte, 0)
		}
		bigByte = append(bigByte, PointBig[z].Bytes()...)
		PointByte = append(PointByte, bigByte...)
	}

	pointS := a.suite.G2().Point()
	err1 := pointS.UnmarshalBinary(PointByte)
	if err1 != nil {
		fmt.Println("501", err1)
	}

	for i := 0; i < len(a.enrollNodes); i++ {
		for j := 0; j < len(PK[i]); j++ {
			// 构造Point累加形式的S
			PKbytes := make([]byte, 0)

			for _, z := range [4]int{1, 0, 3, 2} {
				bigByte := PK[i][j][z].Bytes()
				sub := 32 - len(bigByte)
				bigByte = make([]byte, 0)
				for i := 0; i < sub; i++ {
					bigByte = append(bigByte, 0)
				}

				bigByte = append(bigByte, PK[i][j][z].Bytes()...)
				PKbytes = append(PKbytes, bigByte...)

			}
			PKpoint := a.suite.G2().Point()
			err := PKpoint.UnmarshalBinary(PKbytes)
			if err != nil {
				fmt.Println("translate PK ", err)
			}

			pointS = a.suite.G2().Point().Add(pointS, PKpoint)

			pkSet = append(pkSet, PK[i][j])

		}

	}

	MulSignature := a.suite.G1().Point().Null()

	// MulR := a.suite.G1().Point().Null()
	// MulY := a.suite.G1().Point().Null()   // schnorr

	MulY := a.suite.G2().Point().Null() // bls

	message, _ := encodeValidateResult(txHash, true, typ)

	hash := sha256.New()
	hash.Write(message)

	messageHash := hash.Sum(nil)
	_hash := a.suite.G1().Point().Mul(a.suite.G1().Scalar().SetBytes(messageHash), nil)

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(PK[i]); j++ {

			// schnorr
			// tmpX := PK[i][j][0]
			// tmpY := PK[i][j][1]
			// tmpXByte := tmpX.Bytes()
			// XByte := make([]byte, 32)

			// for k := 31; k >= 0; k-- {
			// 	if len(tmpXByte)-(len(XByte)-k) >= 0 {
			// 		XByte[k] = tmpXByte[len(tmpXByte)-(len(XByte)-k)]
			// 	} else {
			// 		XByte[k] = 0
			// 	}

			// }
			// for k := 0; k < 32; k++ {
			// 	S[k] = XByte[k]
			// }
			// tmpYByte := tmpY.Bytes()

			// YByte := make([]byte, 32)
			// for k := 31; k >= 0; k-- {
			// 	if len(tmpYByte)-(len(YByte)-k) >= 0 {
			// 		YByte[k] = tmpYByte[len(tmpYByte)-(len(YByte)-k)]
			// 	} else {
			// 		YByte[k] = 0
			// 	}

			// }
			// for k := 0; k < 32; k++ {
			// 	S[k+32] = YByte[k]
			// }
			// pkbytes := S[0:64]

			// bls

			PKbytes := make([]byte, 0)

			for _, z := range [4]int{1, 0, 3, 2} {
				bigByte := PK[i][j][z].Bytes()
				sub := 32 - len(bigByte)

				bigByte = make([]byte, 0)

				for i := 0; i < sub; i++ {
					bigByte = append(bigByte, 0)
				}

				bigByte = append(bigByte, PK[i][j][z].Bytes()...)
				PKbytes = append(PKbytes, bigByte...)

			}
			pk := a.suite.G2().Point()
			err := pk.UnmarshalBinary(PKbytes)

			if err != nil {
				fmt.Println("translate pk ", err, len(PKbytes))
			}

			// 累加S
			summaryS := pointS.Clone()
			summaryS = a.suite.G2().Point().Add(summaryS, pk)
			summarySBytes, err := summaryS.MarshalBinary()
			if err != nil {
				log.Errorf("Marshal point err: %v", err)
			}

			hash1 := sha256.New()

			// hash1.Write(S)
			hash1.Write(summarySBytes)
			aI := hash1.Sum(nil)

			aScalar := a.suite.G1().Scalar().SetBytes(aI)

			MulSignature.Add(MulSignature, a.suite.G1().Point().Mul(aScalar, Signatures[i][j]))
			MulY.Add(MulY, a.suite.G2().Point().Mul(aScalar, pk))

		}
	}

	left := a.suite.Pair(_hash, MulY)
	right := a.suite.Pair(MulSignature, a.suite.G2().Point().Base())
	fmt.Println("435 bls", right.Equal(left))
	a.enrollNodes = []int64{}

	// return true, MulSignature, _hash, MulY, nodes, pkSet, nil
	return true, MulSignature, _hash, MulY, nodes, make([][2]*big.Int, 0), nil

}
