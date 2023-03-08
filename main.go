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
	go task("Tarefa 1") //T2 Nova thread
	go task("Tarefa 2") //T3 Nova thread
	task("Tarefa 3")    // T1
}
