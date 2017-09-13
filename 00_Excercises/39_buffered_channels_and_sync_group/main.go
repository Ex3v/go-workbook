package main

import (
	"fmt"
	"sync"
	"time"
)

var channel = make(chan int, 3)
var sg = sync.WaitGroup{}

func main() {
	fmt.Println("\n\nBuffered channels")

	sg.Add(1)
	go worker()

	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Printf("\nLoop write: %d", i)
	}
	close(channel)
	sg.Wait()
}

func worker() {
	for {
		val, opened := <-channel
		if false == opened {
			sg.Done()
		}

		fmt.Printf("\nChannel read: %d", val)
		time.Sleep(1 * time.Second)
	}
}
