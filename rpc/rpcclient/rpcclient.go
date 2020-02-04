package rpcclient

import (
	"bytes"
	"log"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

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
func (c *Client) RawRequest(method string, params interface{}) (response RPCResponse, err error) {
	// Method may not be empty.
	if method == "" {
		return
	}
	// Marshal parameters as "[]" instead of "null" when no parameters are passed.
	if params == nil {
		params = []json.RawMessage{}
	}

	protocol := "http"

	rpcR := rpcRequest{
		"1.0",
		time.Now().UnixNano(),
		method,
		params}

	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	err = jsonEncoder.Encode(rpcR)
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest("POST", protocol + "://"+c.config.Host, payloadBuffer)
	if err != nil {
		return
	}

	httpReq.Header.Add("Content-Type", "application/json;charset=utf-8")

	// Auth ?
	if len(c.config.User) > 0 || len(c.config.Pass) > 0 {
		httpReq.SetBasicAuth(c.config.User, c.config.Pass)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		log.Printf("RawRequest: httpClient %#v", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &response)
	return
}
