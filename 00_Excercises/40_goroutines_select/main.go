package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var channel = make(chan int)
var quit = make(chan bool)
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("\n\nRange channels")

	wg.Add(1)
	go func() {
		for {
			select {
			case c := <-channel:
				fmt.Println(c)
			case <-quit:
				fmt.Println("Quitting!")
				wg.Done()
				break
			}
		}
	}()

	fmt.Println("Stacking...")

	for i := 0; i < 10; i++ {
		stackToChannel(i)
	}

	quit <- true

	wg.Wait()

}

func stackToChannel(a int) {
	channel <- a
	time.Sleep(time.Duration((rand.Int() % 1000) * int(time.Millisecond)))
}
