package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 4, 8, 16}

	r := make([]int, len(slice)-1)
	a := 0

	for i := range r {
		r[i] = a
		a += 1
	}

	fmt.Println(slice)

	fmt.Println(`Indexes`)

	for i := range slice {
		fmt.Println(i)
	}

	fmt.Println(`Indexes and values`)

	for i, val := range slice {
		fmt.Printf("Index: %d, value: %d\n", i, val)
	}

	fmt.Println("\n\nSubslice!")

	fmt.Println(slice[2:4])
}
