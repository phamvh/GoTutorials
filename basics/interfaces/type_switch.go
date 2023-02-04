package main

import "fmt"

// https://go.dev/tour/methods/16

func main() {
	// since checkType takes an arg of an empty interface, we can pass a vale of any type to it.
	checkType(10)      // int
	checkType("hello") // string
	checkType(3.14)    // unknown type
}

/*
*
checkType
https://go.dev/tour/methods/16
Use switch to check the type of a value held by an interface.
It is similar in syntax to type assertion. Just put the keyword *type* in parens instead of a specific type.
*/
func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}
