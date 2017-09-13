package main

import (
	"fmt"
)

func main() {

	out := Pic(32, 32)

	for _, v := range out {
		fmt.Println("")
		for _, u := range v {
			fmt.Printf("%3d", u)
		}
	}

}

func Pic(dx, dy int) [][]uint8 {
	var out [][]uint8

	for y := 0; y < dy; y++ {

		var row []uint8

		for x := 0; x < dx; x++ {
			//val := (x+y)/2
			//val := x*y
			val := x ^ y
			row = append(row, uint8(val))
		}

		out = append(out, row)
	}

	return out
}
