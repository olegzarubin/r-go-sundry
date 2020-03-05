package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func startWorker(in int, wg *sync.WaitGroup) {

	defer wg.Done() //уменьшаем счетчик на 1

	for i := 0; i < iterationsNum; i++ {
		fmt.Printf(formatWork(in, i))
		runtime.Gosched()
	}

}

func main() {

	wg := &sync.WaitGroup{} // инициализируем группу

	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) //добавляем воркер
		go startWorker(i, wg)
	}

	time.Sleep(time.Millisecond)

	wg.Wait() //ожидаем пока wg.Done() приведет счетчик к 0
}

func formatWork(num, it int) string {
	return fmt.Sprintln(strings.Repeat("  ", num), "*",
		strings.Repeat("  ", goroutinesNum-num),
		"worker:", num, "iteration:", it)
}
