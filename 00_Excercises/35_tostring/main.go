package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

type IPAddr [4]byte

func (addr IPAddr) String() string {

	s := []string{}
	for _, v := range addr {
		s = append(s, strconv.Itoa(int(v)))
	}

	return strings.Join(s, ".")
}

func main() {

	p1 := Person{"Johnny Cash", 42}
	p2 := Person{"Mike Sadowsky", 63}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("\n\nIPAddresses\n\n")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
