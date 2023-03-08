package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for value := range data {
		fmt.Printf("Worker %d received %d\n", workerId, value)
		time.Sleep(time.Second)

	}
}

func main() {
	ch := make(chan int)
	qtsWorkers := 100

	// Inicializa ops workers
	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	// joga a carga para os workers
	for i := 0; i < 100000; i++ {
		ch <- i
	}

}
