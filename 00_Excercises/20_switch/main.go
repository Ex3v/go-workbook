package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Print("Go runs on ")

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's saturday?")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow!")
	case today + 2:
		fmt.Println("Day after tomorrow!")
	default:
		fmt.Println("Too far away :(")
	}

}
