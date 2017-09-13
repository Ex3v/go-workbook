package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {

	m := make(map[string]Point)

	m["start"] = Point{x: 0, y: 0}
	m["end"] = Point{x: 255, y: 255}

	fmt.Println(m)
	fmt.Println("only start!", m["start"])

	literalMap := map[string]Point{
		"upper left corner": {-255, 255},
		"lower left corner": {-255, -255},
	}

	literalMap["botton right corner"] = Point{255, -255}

	fmt.Println(literalMap)

	fmt.Println("checking presence of upper right corner")

	_, ok := literalMap["upper right corner"]

	fmt.Println("Existence of upper right corner equals to:", ok)

	delete(literalMap, "upper left corner")
}
