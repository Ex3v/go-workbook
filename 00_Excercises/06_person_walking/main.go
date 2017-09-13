package main

import (
	"fmt"
)

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintf("%s is walking!", p.fname)
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

	s1 := p1.walk()

	fmt.Println(s1)
}
