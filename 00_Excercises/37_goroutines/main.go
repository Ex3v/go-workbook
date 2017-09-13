package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("pomidor")
	fmt.Println("Generating threads")
	generateThreads(10)
	time.Sleep(2 * time.Second)
}

func generateThreads(amount int) {

	for i := 0; i < amount; i++ {
		go threadRunner(i)
	}
}

func threadRunner(threadNumber int) {
	fmt.Println("I'm a thread with number: ", threadNumber)
}
