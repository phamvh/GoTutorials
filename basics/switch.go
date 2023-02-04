package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	no need of break
	switch can be a var, and need not be integer. Can be anything
	*/
	day := 9
	switch day { // no parens needed around "day"
	case 1: // case value must have the same type as switch var.
		fmt.Println("Monday") // no "break" needed
	case 2, 3:
		fmt.Println("Tuesday or Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	case 8, 9, 10: // multiple values can be separated by commas
		fmt.Println("Saturn day")
	default:
		fmt.Println("Mars day")
	}

	// switch with no condition is the same as switch true.
	// It can be a clean way to write long if-then-else chain
	t := time.Now()
	switch { // no var here
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")

	}
}
