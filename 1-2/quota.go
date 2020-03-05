package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {

	quotaCh <- struct{}{} // берем свободный слот
	defer wg.Done()

	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))

		if j%2 == 0 {
			<-quotaCh             // возвращаем слот
			quotaCh <- struct{}{} // берем слот
		}
	}
	<-quotaCh // возвращаем слот

}

func main() {

	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit) // создаем канал квоты

	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}

	time.Sleep(time.Millisecond)
	wg.Wait()

}

func formatWork(num, it int) string {
	return fmt.Sprintln(strings.Repeat("  ", num), "*",
		strings.Repeat("  ", goroutinesNum-num),
		"worker:", num, "iteration:", it)
}
