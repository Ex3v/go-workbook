package main

import (
	"fmt"
)

type gator int

func (g gator) String() string {
	return string(int(g))
}

func (g gator) greeting() {
	fmt.Println("\nHello! I am a gator")
}

func main() {
	var g1 gator
	g1 = 42
	var x int
	x = 5
	fmt.Printf("Value: %s, type: %T", g1, g1)
	fmt.Printf("\nValue: %d, type: %T", x, x)

	x = int(g1)

	fmt.Printf("\nNew X Value: %d, type: %T", x, x)

	g1.greeting()
}
