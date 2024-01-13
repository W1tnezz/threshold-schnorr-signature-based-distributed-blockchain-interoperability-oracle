package node

type Config struct {
	BindAddress string
	PrivateKey  string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	Kafka       KafkaConfig
	Reputation  int64
}

type ContractsConfig struct {
	RegistryContractAddress string
	OracleContractAddress   string
	DistKeyContractAddress  string
}

type EthereumConfig struct {
	TargetAddress string
	SourceAddress string
	PrivateKey    string
	ChainID       int64
}

type KafkaConfig struct {
	IpAddress string
	Topic     string
	Partition int64
}
