package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	fmt.Println("Let's try some interfaces shall we?")

	u := MyFloat(1.23)
	v := Vertex{2, 4}

	fmt.Println(u, v)

	fmt.Println("And the abs's are:")
	fmt.Println(getAbs(u), getAbs(v))
}

func getAbs(a Abser) float64 {
	return a.Abs()
}
