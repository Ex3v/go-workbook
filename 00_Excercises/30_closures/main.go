package main

import (
	"fmt"
)

func counter() func(int) int {

	sum := 0
	return func(i int) int {
		sum += i

		return sum
	}
}

func fibonacci() func() int {
	val1, val2 := 0, 0

	return func() int {

		if val1 == 0 && val2 == 0 {
			val2 = 1
			return 1
		}

		ret := val1 + val2

		val1 = val2
		val2 = ret

		return ret
	}
}

func main() {

	closure := func(a, b int) int { return a + b }

	fmt.Println(closure(1, 2))

	sumator := counter()
	var result int
	for i := 0; i < 10; i++ {
		result = sumator(i)
	}

	fmt.Println(result)

	fmt.Println("Fibonacci!")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("\n%d.\t%d", i+1, f())
	}
}
