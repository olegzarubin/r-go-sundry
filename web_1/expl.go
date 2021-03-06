package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//var baseUrl = "http://127.0.0.1:" + settings.port + "/api/"
var baseURL = "http://localhost:13370/api/"

//var baseUrl = "http://chain.cspn.io" + "" + "/api/"

/*get_difficulty: function(cb) {
  if (settings.use_rpc) {
    rpcCommand([{method:'getdifficulty', parameters: []}], function(response){
      return cb(response);
    });
  } else {
    var uri = base_url + 'getdifficulty';
    request({uri: uri, json: true}, function (error, response, body) {
      return cb(body);
    });
  }
},*/

func getDifficulty() {
	url := baseURL + "getdifficulty"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend (getDifficulty)", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.Get (getDifficulty) body %#v\n\n\n", string(respBody))
}

/*get_connectioncount: function(cb) {
  if (settings.use_rpc) {
    rpcCommand([{method:'getconnectioncount', parameters: []}], function(response){
      return cb(response);
    });
  } else {
    var uri = base_url + 'getconnectioncount';
    request({uri: uri, json: true}, function (error, response, body) {
      return cb(body);
    });
  }
},*/

func getConnectioncount() {
	url := baseURL + "getconnectioncount"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend (getConnectioncount)", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.Get (getConnectioncount) body %#v\n\n\n", string(respBody))
}

/*get_blockcount: function(cb) {
  if (settings.use_rpc) {
    rpcCommand([{method:'getblockcount', parameters: []}], function(response){
      return cb(response);
    })
  } else {
    var uri = base_url + 'getblockcount';
    request({uri: uri, json: true}, function (error, response, body) {
      return cb(body);
    });
  }
},*/

func getBlockcount() {
	url := baseURL + "getblockcount"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend (getBlockcount)", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	log.Printf("http.Get (getBlockcount) body %#v\n\n\n", string(respBody))
}

/*get_maxmoney: function(cb) {
  if (settings.use_rpc) {
    rpcCommand([{method:'getmaxmoney', parameters: []}], function(response){
      return cb(response);
    });
  } else {
    var uri = base_url + 'getmaxmoney';
    request({uri: uri, json: true}, function (error, response, body) {
      return cb(body);
    });
  }
},*/

func getMaxmoney() {
	url := baseURL + "getmaxmoney"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend (getMaxmoney)", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.Get (getMaxmoney) body %#v\n\n\n", string(respBody))
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

	//getDifficulty()
	//getConnectioncount()
	getBlockcount()
	//getMaxmoney()

	//	runGetFullReq()
	//	runTransportAndPost()

}
