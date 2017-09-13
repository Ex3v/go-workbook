package main

import (
	"fmt"
)

func main() {

	fmt.Println("PubSub example\n")

	pub := Pub{}
	pub.registerReceiver(stringReceiver{})
	pub.registerReceiver(boolReceiver{})
	pub.registerReceiver(floatReceiver{})

	a := "Banana"
	b := 3.14
	c := true

	pub.handle(a)
	pub.handle(b)
	pub.handle(c)
}

type Rec interface {
	supports(interface{}) bool
	handle(interface{})
}

type Pub struct {
	receivers []Rec
}

func (pub Pub) handle(item interface{}) {

	for _, receiver := range pub.receivers {
		if receiver.supports(item) {
			receiver.handle(item)
		}
	}
}

func (pub *Pub) registerReceiver(rec Rec) {
	pub.receivers = append(pub.receivers, rec)
}

type stringReceiver struct{}

type floatReceiver struct{}

type boolReceiver struct{}

func (sr stringReceiver) supports(v interface{}) bool {
	_, ok := v.(string)

	return ok
}

func (sr stringReceiver) handle(item interface{}) {
	fmt.Printf("This is surely a string with value '%s'\n", item)
}

func (fr floatReceiver) supports(v interface{}) bool {
	_, ok := v.(float64)

	return ok
}

func (fr floatReceiver) handle(item interface{}) {
	fmt.Printf("I smell float: '%.3f'\n", item)
}

func (br boolReceiver) supports(v interface{}) bool {
	_, ok := v.(bool)

	return ok
}

func (br boolReceiver) handle(item interface{}) {
	fmt.Printf("Bool of value: '%t'\n", item)
}
