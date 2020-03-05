package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {

	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched()
	}
	printFinishWork(workerNum)

}

func main() {

	worketInput := make(chan string, 2)

	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, worketInput)
	}

	months := []string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		worketInput <- monthName
	}
	close(worketInput)

	time.Sleep(time.Millisecond)

}

func formatWork(num int, text string) string {
	return fmt.Sprintln(strings.Repeat("  ", num), "*",
		strings.Repeat("  ", goroutinesNum-num),
		"worker:", num, "data:", text)
}

func printFinishWork(workerNum int) {
	fmt.Println("worker", workerNum, "finished")
}
