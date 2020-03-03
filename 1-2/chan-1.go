package main

import (
	"fmt"
	"runtime"
)

func main() {
	in := make(chan int, 0)

	go func(out chan<- int) {
		for i := 1; i <= 10; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
			runtime.Gosched()
		}
		close(out)
		fmt.Println("generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}

	fmt.Scanln()
}
