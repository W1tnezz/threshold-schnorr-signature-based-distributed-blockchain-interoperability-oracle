package node

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net"

	"go.dedis.ch/kyber/v3/util/random"

	"node/internal/pkg/kyber/pairing/bn256"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnsafeOracleNodeServer
	server          *grpc.Server
	serverLis       net.Listener
	targetEthClient *ethclient.Client
	sourceEthClient *ethclient.Client
	oracleContract  *OracleContractWrapper
	suite           pairing.Suite

	ecdsaPrivateKey   *ecdsa.PrivateKey
	PrivateKey        []kyber.Scalar
	account           common.Address
	connectionManager *ConnectionManager

	validator    *Validator
	aggregator   *Aggregator
	isAggregator bool
	chainId      *big.Int
	reputation   int64
}

func NewOracleNode(c Config) (*OracleNode, error) {
	server := grpc.NewServer()
	serverLis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		return nil, fmt.Errorf("listen on %s: %v", c.BindAddress, err)
	}
	// 创建一个连接以太坊的客户端，TargetAddress是以太坊的目标地址
	targetEthClient, err := ethclient.Dial(c.Ethereum.TargetAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}
	// 这个也是连接以太坊的连接客户端
	sourceEthClient, err := ethclient.Dial(c.Ethereum.SourceAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}
	// 区块链的ID
	chainId := big.NewInt(c.Ethereum.ChainID)

	// 注册

	if err != nil {
		return nil, fmt.Errorf("registry contract: %v", err)
	}

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), targetEthClient)
	oracleContractWrapper := &OracleContractWrapper{
		OracleContract: oracleContract,
	}
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("dist key contract: %v", err)
	}

	suite := bn256.NewSuite()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}
	// schnorrPrivateKey := make([]kyber.Scalar, 0)
	privateKey := make([]kyber.Scalar, 0)

	reputation := int64(c.Reputation)

	// for i := int64(0); i < reputation; i++ {
	// 	schnorrPrivateKey = append(schnorrPrivateKey, suite.G1().Scalar().Pick(random.New()))
	// }
	for i := int64(0); i < reputation; i++ {
		privateKey = append(privateKey, suite.G2().Scalar().Pick(random.New()))
	}
	if err != nil {
		return nil, fmt.Errorf("hex to scalar: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	connectionManager := NewConnectionManager(oracleContractWrapper, account)
	RAll := make(map[common.Address]kyber.Point)
	enrollNodes := []int64{}

	// 初始化kafka Writer 和 Reader
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(c.Kafka.IpAddress),
		Topic:                  c.Kafka.Topic,
		RequiredAcks:           kafka.RequireAll,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		Async:                  true,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{c.Kafka.IpAddress},
		Topic:     c.Kafka.Topic,
		Partition: int(c.Kafka.Partition),
		MaxBytes:  10e6, // 10MB
	})

	validator := NewValidator(
		suite,
		oracleContractWrapper,
		ecdsaPrivateKey,
		sourceEthClient,
		connectionManager,
		RAll,
		account,
		writer,
		reader,

		privateKey, // 私钥
		reputation,
	)
	aggregator := NewAggregator(
		suite,
		targetEthClient,
		connectionManager,
		oracleContractWrapper,
		account,
		ecdsaPrivateKey,
		chainId,
		enrollNodes,
	)
	node := &OracleNode{
		server:          server,
		serverLis:       serverLis,
		targetEthClient: targetEthClient,
		sourceEthClient: sourceEthClient,
		oracleContract:  oracleContractWrapper,
		suite:           suite,

		ecdsaPrivateKey:   ecdsaPrivateKey,
		PrivateKey:        privateKey,
		account:           account,
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
		isAggregator:      false,
		chainId:           chainId,
		reputation:        reputation,
	}

	RegisterOracleNodeServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {
	// 创建连接
	if err := n.connectionManager.InitConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	go func() {
		if err := n.validator.ListenAndProcess(n); err != nil {
			log.Errorf("Watch and handle DKG log: %v", err)
		}
	}()

	go func() {
		if err := n.connectionManager.WatchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("Watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleValidationRequestsLog(context.Background(), n); err != nil {
			log.Errorf("Watch and handle ValidationRequest log: %v", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {
	isRegistered, err := n.oracleContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	schnorrPublicKey := make([]kyber.Point, 0)
	for _, schnorrPrivateKey := range n.PrivateKey {
		schnorrPublicKey = append(schnorrPublicKey, n.suite.G1().Point().Mul(schnorrPrivateKey, nil))
	}

	blsPublicKey := make([]kyber.Point, 0)
	for _, privateKey := range n.PrivateKey {

		blsPublicKey = append(blsPublicKey, n.suite.G2().Point().Mul(privateKey, nil))
	}

	bSchnorr := make([][2]*big.Int, 0) // schnorr
	bBls := make([][4]*big.Int, 0)     // bls

	for _, publicKey := range schnorrPublicKey {
		publicKeyToBig, err := G1PointToBig(publicKey)

		if err != nil {
			return fmt.Errorf("marshal public key: %v", err)
		}
		bSchnorr = append(bSchnorr, publicKeyToBig)
		log.Info(len(publicKeyToBig[0].Bytes()), " ", len(publicKeyToBig[1].Bytes()))
	}

	for _, publicKey := range blsPublicKey {
		// publicKeyByte, _ := publicKey.MarshalBinary()

		// p := n.suite.G2().Point()
		// p.UnmarshalBinary(publicKeyByte)
		// fmt.Println("246", p.Equal(publicKey), p, publicKey)

		publicKeyToBig, err := G2PointToBig(publicKey)
		// PKbytes := make([]byte, 0)

		// for _, z := range [4]int{1, 0, 3, 2} {
		// 	bigByte := publicKeyToBig[z].Bytes()
		// 	sub := 32 - len(bigByte)

		// 	bigByte = make([]byte, 0)

		// 	for i := 0; i < sub; i++ {
		// 		bigByte = append(bigByte, 0)
		// 	}

		// 	bigByte = append(bigByte, publicKeyToBig[z].Bytes()...)
		// 	PKbytes = append(PKbytes, bigByte...)

		// }
		// for i, _ := range publicKeyByte {
		// 	fmt.Println(publicKeyByte[i] == PKbytes[i])
		// }
		// fmt.Println(publicKeyByte, PKbytes)

		// publicKeyToByte := make([]byte, 0)

		// for z := 0; z < 4; z++ {
		// 	bigByte := publicKeyToBig[z].Bytes()
		// 	for _, b := range bigByte {
		// 		publicKeyToByte = append(publicKeyToByte, b)
		// 	}

		// }
		// fmt.Println("247", len(publicKeyToByte), publicKeyToByte)
		// pointS := n.suite.G2().Point().Null()
		// pointS.UnmarshalBinary(publicKeyToByte)

		// fmt.Println("pointS", pointS)

		// re1 , _ := ScalarToBig(aScalarG1)
		// re2 , _ := ScalarToBig(aScalarG2)

		if err != nil {
			return fmt.Errorf("marshal public key: %v", err)
		}
		bBls = append(bBls, publicKeyToBig)
	}

	minStake, err := n.oracleContract.MINSTAKE(nil)
	if err != nil {
		return fmt.Errorf("min stake: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}
	reputation := big.NewInt(n.reputation)
	auth.Value = minStake.Mul(minStake, reputation)

	if !isRegistered {

		_, err = n.oracleContract.RegisterOracleNode(auth, ipAddr, bSchnorr, big.NewInt(n.reputation))
		if err != nil {
			return fmt.Errorf("register iop node: %w", err)
		}
	}
	return nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()
	n.targetEthClient.Close()
	n.sourceEthClient.Close()
	n.connectionManager.Close()
}
