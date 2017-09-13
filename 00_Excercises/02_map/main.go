package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"age": 19,
		"times had pinky hit":  155,
		"usual sleep duration": 6,
	}

	fmt.Println(m)

	fmt.Println("\n\nPrinting just keys:")

	for i := range m {
		fmt.Println(i)
	}

	fmt.Println("\n\nPrinting keys and values:")

	for i, val := range m {
		fmt.Printf("%s: %d\n", i, val)
	}
}
