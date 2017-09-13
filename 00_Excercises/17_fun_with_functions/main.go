package main

import (
	"fmt"
	"math"
)

func foo(a int, b int) (v, u int) { fmt.Println("potato"); return b, a }

func main() {
	a, b := foo(5, 8)

	x := func(a, b int) (r int) { r = int(math.Pow(float64(a), float64(b))); return }(2, 4)
	fmt.Println(a, b, x)

	var min, max, step int = 0, 10, 2

	for i := min; i <= max; i += step {
		fmt.Println(i)
	}
}
