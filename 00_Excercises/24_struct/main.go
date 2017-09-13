package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {

	v := Vertex{Y: 3, X: 2}
	p := &v

	p.X = 8
	v.Y = 4

	fmt.Println(v, "\n\n", p)
	fmt.Printf("%T %T", v, p)
}
