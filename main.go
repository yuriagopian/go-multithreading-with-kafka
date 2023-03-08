package main

import (
	"fmt"
	"time"
)

func worker(workerId, data chan int) {
	for value := range data {
		fmt.Printf("Worker %d received %d\n", workerId, value)
		time.Sleep(time.Second)

	}
}

func main() {
	ch := make(chan int)
	qtsWorkers := 3

	for i := 0; i < qtsWorkers; i++ {
		ch <- i
	}

}
