package main

import (
	"fmt"
)

func main() {

	fmt.Println("Arrays stuff")

	var a [3]string

	a[0] = "potato"
	a[1] = "banana"

	fmt.Println(a, len(a))
	fmt.Println("\n\nSlices stuff")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	bananas := []string{"one", "two", "three"}
	bananas = append(bananas, "I love oreos")
	bananas = append(bananas, "And pens")
	bananas = append(bananas, "And pineapples")
	bananas = append(bananas, "And pineapplepens")

	var s []int = primes[1:4]

	fmt.Println(s, bananas, len(bananas), cap(bananas))

	var u []int
	fmt.Println(u, len(u), cap(u))
	if u == nil {
		fmt.Println("nil!")
	}

}
