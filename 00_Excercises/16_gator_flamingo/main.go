package main

import (
	"fmt"
	"math/rand"
)

type gator int

type flamingo bool

func (g gator) String() string {
	return string(int(g))
}

func (g gator) greeting() {
	fmt.Println("Hello! I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello! I am a flamingo")
}

type swampCreature interface {
	greeting()
}

func main() {
	var g1 gator
	var f1 flamingo
	g1 = 42
	f1 = false
	var x int
	x = 5
	fmt.Printf("Value: %s, type: %T", g1, g1)
	fmt.Printf("\nValue: %d, type: %T", x, x)

	x = int(g1)

	fmt.Printf("\nNew X Value: %d, type: %T\n", x, x)

	bayou(g1)
	bayou(f1)

	fmt.Println(rand.Intn(51))
}

func bayou(s swampCreature) {
	s.greeting()
}
