package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var totalOperations int32 = 0

func inc() {
	time.Sleep(time.Duration(rand.Intn(100)+10) * time.Millisecond)
	//totalOperations++
	atomic.AddInt32(&totalOperations, 1)
}

func main() {

	for i := 0; i < 1000; i++ {
		go inc()
	}

	time.Sleep(1 * time.Second)

	// ожидается 1000 но по факту будет меньше
	fmt.Println("total operation = ", totalOperations)
}
