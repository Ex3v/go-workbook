package main

import (
	"fmt"
)

type gator int

func main() {
	var g1 gator
	g1 = 42
	var x int
	x = 5
	fmt.Printf("Value: %s, type: %T", g1, g1)
	fmt.Printf("\nValue: %d, type: %T", x, x)
}
