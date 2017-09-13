package main

import (
	"fmt"
)

type vehicle struct {
	color string
	doors int
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() {
	fmt.Println("I move small stuff")
}

func (t truck) transportationDevice() {
	fmt.Println("I move big stuff")
}

type transportation interface {
	transportationDevice()
}

func main() {

	t1 := truck{
		vehicle{
			"black",
			2,
		},
		true,
	}

	s1 := sedan{
		vehicle{
			"red",
			4,
		},
		false,
	}

	report(s1)
	report(t1)
}

func report(t transportation) {
	t.transportationDevice()
}
