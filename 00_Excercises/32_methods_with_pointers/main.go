package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(multiplier int) {
	v.Y *= float64(multiplier)
	v.X *= float64(multiplier)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	fmt.Println("Pointers TEST")

	v := Vertex{1, 2}

	fmt.Println(v)

	fmt.Println("Now the magic happens")

	v.Scale(3)

	fmt.Println(v)

	u := Vertex{2, 4}
	fmt.Println("Now time for some functions")
	fmt.Println(u, Abs(u))
	fmt.Println("Let's roll")

	Scale(&u, 3)
	fmt.Println(u, Abs(u))

}
