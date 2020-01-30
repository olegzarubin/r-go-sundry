package main

import (
	"log"

	"github.com/toorop/go-bitcoind"
)

const (
	serverHost       = "localhost"
	serverPort       = 13371
	user             = "csportsrpc"
	pass             = "V6e6cPTTeneTTdnosLLD3cUGnfh67gdmRTGvA5YB3UZh"
	useSSL           = false
	walletPassphrase = ""
)

func main() {

	client, err := bitcoind.New(serverHost, serverPort, user, pass, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	//	defer client.Shutdown()	??

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
