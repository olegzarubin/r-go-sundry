package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"strconv"
	"sync/atomic"
	"time"
)

type job func(in, out chan interface{})

const (
	MaxInputDataLen = 100
)

var (
	dataSignerOverheat uint32 = 0
	DataSignerSalt            = ""
)

var OverheatLock = func() {
	for {
		if swapped := atomic.CompareAndSwapUint32(&dataSignerOverheat, 0, 1); !swapped {
			fmt.Println("OverheatLock happend")
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}

var OverheatUnlock = func() {
	for {
		if swapped := atomic.CompareAndSwapUint32(&dataSignerOverheat, 1, 0); !swapped {
			fmt.Println("OverheatUnlock happend")
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}

var DataSignerMd5 = func(data string) string {
	OverheatLock()
	defer OverheatUnlock()
	data += DataSignerSalt
	dataHash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	time.Sleep(10 * time.Millisecond)
	return dataHash
}

var DataSignerCrc32 = func(data string) string {
	data += DataSignerSalt
	crcH := crc32.ChecksumIEEE([]byte(data))
	dataHash := strconv.FormatUint(uint64(crcH), 10)
	time.Sleep(time.Second)
	return dataHash
}

func SingleHash(in, out chan interface{}) {

	dataRaw := <-in

	data, ok := dataRaw.(int)
	if !ok {
		fmt.Errorf("cant convert result data to int")
	}
	fmt.Println(data, "SingleHash data", data)

	resMD5 := DataSignerMd5(strconv.Itoa(data))
	fmt.Println(data, "SingleHash md5(data)", resMD5)

	resCRC32MD5 := DataSignerCrc32(resMD5)
	fmt.Println(data, "SingleHash crc32(md5(data))", resCRC32MD5)

	resCRC32 := DataSignerCrc32(strconv.Itoa(data))
	fmt.Println(data, "SingleHash crc32(data)", resCRC32)

	result := resCRC32 + "~" + resCRC32MD5
	fmt.Println(data, "SingleHash result", result)

	out <- result
}

func MultiHash(in, out chan interface{}) {

	dataRaw := <-in

	data, ok := dataRaw.(string)
	if !ok {
		fmt.Errorf("cant convert result data to string")
	}
	fmt.Println(data, "MultiHash data", data)

	result := ""

	for i := 0; i <= 5; i++ {
		res := DataSignerCrc32(strconv.Itoa(i) + data)
		result = result + res
		fmt.Println(data, "MultiHash: crc32(th+step1))", i, res)
	}

	fmt.Println(data, "MultiHash result:", result, "\n")

	out <- result
}

func main() {

	inCh := make(chan interface{})
	outCh := make(chan interface{})

	go SingleHash(inCh, outCh)
	inCh <- 0
	result := <-outCh

	go MultiHash(inCh, outCh)
	inCh <- result
	<- outCh

	// 2 проход

	go SingleHash(inCh, outCh)
	inCh <- 1
	result = <-outCh

	go MultiHash(inCh, outCh)
//	inCh = <-outCh
	inCh <- result
	result = <-outCh

	fmt.Printf("result %#v\n", result)
}
