package main

import (
	"fmt"
	"math"
	"sync"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

var channel = make(chan int)
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("\n\nRange channels")

	fib := []int{1, 2, 3, 4, 5, 8}

	wg.Add(1)
	go walker()
	tree := New(fib)

	tree.Walk(channel)

	close(channel)
}

func walker() {
	for {
		c, ok := <-channel

		if false == ok {
			return
			wg.Done()
		}

		fmt.Println(c)
	}
}

func New(values []int) Tree {
	valLen := len(values)

	if 0 == valLen {
		return Tree{}
	}

	if 1 == valLen {
		return Tree{nil, values[0], nil}
	}

	idx := int(math.Ceil(float64(valLen / 2)))
	leftSlice := values[0:idx]
	val := values[idx]
	rightSlice := values[idx+1:]

	var lTree Tree
	var rTree Tree

	if 0 < len(leftSlice) {
		lTree = New(leftSlice)
	}

	if 0 < len(rightSlice) {
		rTree = New(rightSlice)
	}

	return Tree{&lTree, val, &rTree}
}

func (t *Tree) Walk(c chan int) {
	c <- t.Value

	if nil != t.Left {
		t.Left.Walk(c)
	}

	if nil != t.Right {
		t.Right.Walk(c)
	}
}
