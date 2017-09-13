package main

import "fmt"

type CrazyFloat float64

func (f CrazyFloat) Abs() float64 {
	if f < 0 {
		return float64(f * -1)
	}

	return float64(f)
}

func main() {

	fmt.Println("CrazyFloat tests")

	var f, g CrazyFloat
	f = -1.23
	g = 2.48

	fmt.Println(f.Abs(), g.Abs())
}

// 0 0
// 0 1
//
