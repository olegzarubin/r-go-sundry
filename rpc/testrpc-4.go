package main

import (
	"encoding/json"
	"log"
	"strconv"
	"./rpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host: "localhost:13371",
		User: "csportsrpc",
		Pass: "V6e6cPTTeneTTdnosLLD3cUGnfh67gdmRTGvA5YB3UZh",
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg)
	if err != nil {
		log.Fatal(err)
	}
	//	defer client.Shutdown()

	// Get the miscellaneous info regarding the RPC server.
	r, err := client.RawRequest("getinfo", nil)
	if err != nil {
		log.Printf("Error in RawRequest")
		return
	}

	i := &rpcclient.InfoResult{}
	err = json.Unmarshal(r.Result, &i)

	log.Printf("Getinfo: %#v", i)


	// Get the current block count.
	r, err = client.RawRequest("getblockcount", nil)
	if err != nil {
		log.Fatal(err)
	}

	blockCount, err := strconv.ParseInt(string(r.Result), 10, 32)

	log.Printf("Block count: %d", blockCount)

	// Get the current connection count.
	r, err = client.RawRequest("getconnectioncount", nil)
	if err != nil {
		log.Fatal(err)
	}

	connectionCount, err := strconv.ParseInt(string(r.Result), 10, 32)

	log.Printf("Connection count: %d", connectionCount)

	// Get the current difficulty.
	r, err = client.RawRequest("getdifficulty", nil)
	if err != nil {
		log.Fatal(err)
	}

	difficulty, err := strconv.ParseFloat(string(r.Result), 64)

	log.Printf("Difficulty: %v", difficulty)

}
