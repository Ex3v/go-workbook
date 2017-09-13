package main

import (
	"fmt"
)

type NegativeAgeError int

type Person struct {
	Name string
	Age  int
}

func (e NegativeAgeError) Error() string {
	return fmt.Sprintf("Cannot set age lower than 1! %d given", e)
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func (p *Person) UpdateAge(age int) error {
	if 1 > age {
		return NegativeAgeError(age)
	}

	p.Age = age
	return nil
}

func main() {

	p1 := Person{"Johnny Cash", 42}
	p2 := Person{"Mike Sadowsky", 63}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("This will pass")
	p1.UpdateAge(3)

	fmt.Println(p1)

	fmt.Println("This should not pass")
	e := p1.UpdateAge(-8)

	fmt.Println(p1, e)
}
