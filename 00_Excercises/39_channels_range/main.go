package main

import (
	"fmt"
	"sync"
	"time"
)

var channel = make(chan int)
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("\n\nRange channels")

	go generator()
	wg.Add(2)
	go reader(1)
	go reader(2)

	wg.Wait()
}

func reader(pid int) {
	for c := range channel {
		fmt.Printf("\nReader %d value %d", pid, c)
		time.Sleep(300 * time.Millisecond)
	}
	wg.Done()
}

func generator() {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)
}
