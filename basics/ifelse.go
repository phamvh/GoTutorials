package main

import "fmt"

/*
*
Ifelse is very much like in python.
It seems like all these new languages (scala, go, etc.) are created to mimic python for its simplicity in syntax.
*/
func main() {
	if 10 < 20 { //parens are optional
		fmt.Println("hey")
	}
	if false {
		fmt.Println("Duh")
	}
	if 10 < 20 && 20 < 30 {
		fmt.Println("Sunny")
	} else {
		fmt.Println("rainy")
	}

	x := 1
	if (x > 1) { // parens are optional
		fmt.Println("greater than 1")
	} else if x < 1 {
		fmt.Println("less than 1")
	} else {
		fmt.Println("it is 1")
	}
}
