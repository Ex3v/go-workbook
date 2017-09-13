package main

import (
	"fmt"
)

func main() {

	i := 10

	var j, k *int

	j = &i
	k = &i

	fmt.Printf("\nBefore change: i: %d j: %d k: %d", i, *j, *k)

	i = 30

	fmt.Printf("\n\n first change: i: %d j: %d k: %d", i, *j, *k)

	i = 40
	k = j

	fmt.Printf("\n\nAfter second change: i: %d j: %d k: %d", i, *j, *k)

}
