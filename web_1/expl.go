package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func runGet() {
	//url := "http://127.0.0.1:13370/api/getblockcount"
	url := "http://chain.cspn.io/api/getinfo"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		//log.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
	//log.Printf("http.Get body %#v\n\n\n", string(respBody))
}

func runGetFullReq() {

	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"coursera/golang"},
		},
	}

	//req.URL, _ = url.Parse("http://127.0.0.1:13370/api/getblockcount")
	req.URL, _ = url.Parse("http://chain.cspn.io/api/getinfo")
	//req.URL.Query().Set("user", "rvasily")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		//log.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("runGetFullReq resp %#v\n\n\n", string(respBody))
	//log.Printf("runGetFullReq resp %#v\n\n\n", string(respBody))
}

func runTransportAndPost() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	//data := `{"id": 42, "user": "rvasily"}`
	data := `{}`
	body := bytes.NewBufferString(data)

	//url := "http://127.0.0.1:8080/raw_body"
	url := "http://127.0.0.1:13370/api/getblockcount"
	//url := "http://chain.cspn.io/api/getinfo"
	//req, _ := http.NewRequest(http.MethodPost, url, body)
	req, _ := http.NewRequest(http.MethodGet, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		//log.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("runTransport %#v\n\n\n", string(respBody))
	//log.Printf("runTransport %#v\n\n\n", string(respBody))
}

func main() {

	runGet()
	runGetFullReq()
	runTransportAndPost()

}
