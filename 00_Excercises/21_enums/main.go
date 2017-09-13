package main

import (
	"fmt"
	"reflect"
)

type Color int

const (
	COLOR_RED Color = iota
	COLOR_BLUE
	COLOR_WHITE
)

func main() {

	v := []Color{COLOR_RED, COLOR_BLUE, COLOR_WHITE}
	for i := range v {
		fmt.Println(i)
		fmt.Println(i)
		foo(COLOR_RED)
	}

}

func foo(c Color) {
	fmt.Println(reflect.TypeOf(c))
}
