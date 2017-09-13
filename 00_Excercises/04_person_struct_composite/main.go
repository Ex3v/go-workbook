package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func main() {

	p1 := person{"John", "Black"}

	fmt.Println(p1)
	fmt.Println(p1.lname)
}
