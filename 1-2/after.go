package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World")
}

func main() {

	timer := time.AfterFunc(5*time.Second, sayHello)

	fmt.Scanln()
	timer.Stop()

	//fmt.Scanln()
}
