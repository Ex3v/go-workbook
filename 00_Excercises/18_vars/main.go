package main

import (
	"fmt"
)

const BANANA = "there are bananas"
const (
	BIG   = 1 << 62
	SMALL = BIG >> 19
)

func main() {
	i := 4.0

	fmt.Printf("%.2f, %T, %s", i, i, BANANA)
	fmt.Printf("\n%d\n%d", BIG, SMALL)
}
