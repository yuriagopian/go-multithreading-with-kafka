package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan string)

	// thread 2
	go func() {
		channel <- "Olá Mundo!"
	}()

	// T1
	msg := <-channel
	fmt.Println(msg)

	// go task("Tarefa 1") //T2 Nova thread
	// go task("Tarefa 2") //T3 Nova thread
	// task("Tarefa 3")    // T1
}

func publish(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
}
