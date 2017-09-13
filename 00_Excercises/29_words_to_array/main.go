package main

import (
	"fmt"
	"strings"
)

func main() {

	someString := "This is some string to be splitted to multiple 	words   and so on."

	exploded := strings.Fields(someString)

	counts := map[string]int{}

	for _, word := range exploded {
		_, exists := counts[word]

		if false == exists {
			counts[word] = 0
		}

		counts[word] += 1
	}
}
