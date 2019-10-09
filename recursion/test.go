package main

import "fmt"

func test(n int) int {
	if n == 0 {
		return 0
	}

	fmt.Println("Test", n)

	test(n - 1)

	return 0
}

func main() {

	test(3)
	
}
