package main

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:13371",
		User:         "csportsrpc",
		Pass:         "V6e6cPTTeneTTdnosLLD3cUGnfh67gdmRTGvA5YB3UZh",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

	// Get the current connection count.
	connectionCount, err := client.GetConnectionCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connection count: %d", connectionCount)

	// Get the current difficulty.
	currentDifficulty, err := client.GetDifficulty()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current difficulty: %v", currentDifficulty)

	// Get the miscellaneous info regarding the RPC server.
	infoResult, err := client.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Miscellaneous info: %#v", infoResult)

}
