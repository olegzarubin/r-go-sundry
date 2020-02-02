package rpcclient

import (
	"log"
//	"bytes"
//	"crypto/tls"
//	"encoding/json"
	"errors"
//	"fmt"
//	"io/ioutil"
	"net/http"
//	"time"
)

type ConnConfig struct {
	// Host is the IP address and port of the RPC server you want to connect to.
	Host string
	// User is the username to use to authenticate to the RPC server.
	User string
	// Pass is the passphrase to use to authenticate to the RPC server.
	Pass string
}

type Client struct {
	// config holds the connection configuration assoiated with this client.
	config *ConnConfig
	// httpClient is the underlying HTTP client to use when running in HTTP POST mode.
	httpClient *http.Client
}

func New(config *ConnConfig) (*Client, error) {
	if len(config.Host) == 0 {
		err := errors.New("Bad call missing argument host")
		return nil, err
	}

	httpClient := &http.Client{}

	client := &Client{
		config:       config,
		httpClient:   httpClient,
	}

	log.Printf("Established connection to RPC server %s", config.Host)

	return client, nil
}
