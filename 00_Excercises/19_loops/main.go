package main

import (
	"fmt"
	"math"
)

func step(z, x float64) (u float64) {
	u = z - ((math.Pow(z, 2) - x) / (2 * z))
	return
}

func Sqrt(x float64) (res float64) {

	res = 1
	for i := 1; i < 10; i++ {
		res = step(res, x)
		fmt.Println(res)
	}

	return
}

func main() {
	for i := 0; i <= 100; i += 2 {
		fmt.Print(i, ".")
	}

	a, u := false, 50

	for a == false {
		u--
		if u == 0 {
			a = true
		}
	}

	fmt.Println("\n\npotato")

	x := Sqrt(256)

	fmt.Println(x)
}
