package node

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net"

	"node/internal/pkg/kyber/pairing/bn256"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	"log" 
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

	ecdsaPrivateKey *ecdsa.PrivateKey
	PrivateKey      kyber.Scalar
	account         common.Address

	validator    *Validator
	aggregator   *Aggregator
	isAggregator bool
	chainId      *big.Int
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

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	registryContract, err := NewRegistry(common.HexToAddress(c.Contracts.RegistryContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("registry contract: %v", err)
	}

	dkgContract, err := NewDKG(common.HexToAddress(c.Contracts.DistKeyContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("dkg contract: %v", err)
	}
	oracleContractWrapper := &OracleContractWrapper{
		Registry:       registryContract,
		DKG:            dkgContract,
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
	privateKey := suite.G1().Scalar().Pick(suite.RandomStream())

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	RAll := make(map[common.Address]kyber.Point)

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

		RAll,
		account,
		writer,
		reader,

		privateKey, // 私钥
		false,
	)
	aggregator := NewAggregator(
		suite,
		targetEthClient,

		oracleContractWrapper,
		account,
		ecdsaPrivateKey,
		chainId,
		false,
	)
	node := &OracleNode{
		server:          server,
		serverLis:       serverLis,
		targetEthClient: targetEthClient,
		sourceEthClient: sourceEthClient,
		oracleContract:  oracleContractWrapper,
		suite:           suite,

		ecdsaPrivateKey: ecdsaPrivateKey,
		PrivateKey:      privateKey,
		account:         account,
		validator:       validator,
		aggregator:      aggregator,
		isAggregator:    false,
		chainId:         chainId,
	}

	RegisterOracleNodeServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {

	go func() {
		if err := n.validator.ListenAndProcess(n); err != nil {
			log.Println("Watch and handle DKG log: %v", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleValidationRequestsLog(context.Background(), n); err != nil {
			log.Println("Watch and handle ValidationRequest log: ", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleDKGLog(context.Background()); err != nil {
			log.Println("Watch and handle ValidationRequest log: ", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}
	minStack, err := n.oracleContract.MinStake(nil)
	auth.Value = minStack

	publicKey := n.suite.G1().Point().Mul(n.PrivateKey, nil)
	publickeyBigint, err := G1PointToBig(publicKey)
	if err != nil {
		return fmt.Errorf("scalarToBig err : %w", err)
	}
	_, err = n.oracleContract.Register(auth, ipAddr, publickeyBigint)
	if err != nil {
		return fmt.Errorf("register iop node: %w", err)
	}

	return nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()
	n.targetEthClient.Close()
	n.sourceEthClient.Close()
}
