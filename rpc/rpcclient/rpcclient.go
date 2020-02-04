package rpcclient

import (
	"bytes"
	"log"

	//	"crypto/tls"
	"encoding/json"
	"errors"

	//	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// A ConnConfig represent a connection config
type ConnConfig struct {
	// Host is the IP address and port of the RPC server you want to connect to.
	Host string
	// User is the username to use to authenticate to the RPC server.
	User string
	// Pass is the passphrase to use to authenticate to the RPC server.
	Pass string
}

// A Client represents a JSON RPC client (over HTTP)
type Client struct {
	// config holds the connection configuration assoiated with this client.
	config *ConnConfig
	// httpClient is the underlying HTTP client to use when running in HTTP POST mode.
	httpClient *http.Client
}

// rpcRequest represent a RCP request
type rpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// A RPCResponce represent a RCP responce
type RPCResponse struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result"`
	Err    *RPCError       `json:"error"`
}

// RPCErrorCode represents an error code to be used as a part of an RPCError
// which is in turn used in a JSON-RPC Response object.
//
// A specific type is used to help ensure the wrong errors aren't used.
type RPCErrorCode int

// RPCError represents an error that is used as a part of a JSON-RPC Response
// object.
type RPCError struct {
	Code    RPCErrorCode `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

// An Info represent a response to getmininginfo
type Info struct {
	// The server version
	Version uint32 `json:"version"`

	// The protocol version
	Protocolversion uint32 `json:"protocolversion"`

	// The wallet version
	Walletversion uint32 `json:"walletversion"`

	// The total bitcoin balance of the wallet
	Balance float64 `json:"balance"`

	// The current number of blocks processed in the server
	Blocks uint32 `json:"blocks"`

	// The time offset
	Timeoffset int32 `json:"timeoffset"`

	// The number of connections
	Connections uint32 `json:"connections"`

	// Tthe proxy used by the server
	Proxy string `json:"proxy,omitempty"`

	// Tthe current difficulty
	Difficulty float64 `json:"difficulty"`

	// If the server is using testnet or not
	Testnet bool `json:"testnet"`

	// The timestamp (seconds since GMT epoch) of the oldest pre-generated key in the key pool
	Keypoololdest uint64 `json:"keypoololdest"`

	// How many new keys are pre-generated
	KeypoolSize uint32 `json:"keypoolsize,omitempty"`

	// The timestamp in seconds since epoch (midnight Jan 1 1970 GMT) that the wallet is unlocked for transfers, or 0 if the wallet is locked
	UnlockedUntil int64 `json:"unlocked_until,omitempty"`

	// the transaction fee set in btc/kb
	Paytxfee float64 `json:"paytxfee"`

	// Minimum relay fee for non-free transactions in btc/kb
	Relayfee float64 `json:"relayfee"`

	//  Any error messages
	Errors string `json:"errors"`
}

// New return a new Client
func New(config *ConnConfig) (*Client, error) {
	if len(config.Host) == 0 {
		err := errors.New("Bad call missing argument host")
		return nil, err
	}

	httpClient := &http.Client{}

	client := &Client{
		config:     config,
		httpClient: httpClient,
	}

	log.Printf("Established connection to RPC server %s", config.Host)

	return client, nil
}

// RawRequest call prepare & exec the request
func (c *Client) RawRequest(method string, params interface{}) (rr RPCResponse, err error) {
	// Method may not be empty.
	if method == "" {
		return
	}
	// Marshal parameters as "[]" instead of "null" when no parameters are passed.
	if params == nil {
		params = []json.RawMessage{}
	}

	log.Printf("RawRequest: %#v", 1)

	//	connectTimer := time.NewTimer(time.Duration(c.timeout) * time.Second)
	rpcR := rpcRequest{
		"1.0",
		time.Now().UnixNano(),
		method,
		params}

	log.Printf("RawRequest: %#v", rpcR)
	payloadBuffer := &bytes.Buffer{}

	jsonEncoder := json.NewEncoder(payloadBuffer)
	//	jsonEncoder, err := json.Marshal(payloadBuffer)
	log.Printf("jsonEncoder: %#v", jsonEncoder)

	err = jsonEncoder.Encode(rpcR)
	if err != nil {
		return
	}

	log.Printf("RawRequest: %#v", 2)

	req, err := http.NewRequest("POST", "http://"+c.config.Host, payloadBuffer)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	log.Printf("RawRequest: %#v", 3)

	// Auth ?
	if len(c.config.User) > 0 || len(c.config.Pass) > 0 {
		req.SetBasicAuth(c.config.User, c.config.Pass)
	}

	log.Printf("RawRequest: %#v", 4)
	//	log.Printf("RawRequest: %#v", req)

	//	resp, err := c.doTimeoutRequest(connectTimer, req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("RawRequest: httpClient %#v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("RawRequest: %#v", 5)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Printf("RawRequest: %#v", 6)

	err = json.Unmarshal(data, &rr)
	return
}
