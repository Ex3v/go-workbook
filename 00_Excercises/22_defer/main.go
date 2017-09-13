package main

import (
	"fmt"
	"time"
)

func main() {
	wrapper(2)

	fmt.Println("\n\n now comes defer stack!")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}

func wrapper(sleep int) {
	defer deferred()
	fmt.Printf("Defer called and now I wait %d seconds", sleep)
	time.Sleep(time.Duration(sleep) * time.Second)
}

func deferred() {
	fmt.Println("\n\nI am deferred")
}
