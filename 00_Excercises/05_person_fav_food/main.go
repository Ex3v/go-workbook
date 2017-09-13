package main

import (
	"fmt"
)

type person struct {
	fname   string
	lname   string
	favFood []string
}

func main() {

	p1 := person{
		"John",
		"Black",
		[]string{
			"potato",
			"banana", "peanut butter",
		},
	}

	fmt.Println(p1)

	fmt.Println("\nListing only fav foods!\n")

	for _, val := range p1.favFood {
		fmt.Printf("Course: %s\n", val)
	}
}
