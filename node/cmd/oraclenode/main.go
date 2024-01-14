package main

import (
	"flag"
	"log"
	"node/pkg/node"
	"os"
	"os/signal"

	"github.com/spf13/viper"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	commonConfig := "./configs/common.json"
	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	var config node.Config

	// 公共配置
	viper.SetConfigFile(*&commonConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config: %v", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unmarshal common config into struct, %v", err)
	}

	// 节点配置
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config: %v", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unmarshal config into struct, %v", err)
	}

	log.Println("Loaded config file ", *configFile)

	node, err := node.NewOracleNode(config) // 根据config初始化node
	if err != nil {
		log.Fatalf("New oracle node: %v", err)
	}

	go func() {
		if err := node.Run(); err != nil {
			log.Fatalf("Run node: %v", err)
		} // 运行node
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	node.Stop()
}
