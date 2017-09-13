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
	fmt.Println(`Creating vehicles`)
}
