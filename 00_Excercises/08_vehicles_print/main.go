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

	fmt.Println("Truck:")
	fmt.Println(t1)

	fmt.Println("Sedan:")
	fmt.Println(s1)
}
