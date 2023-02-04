package main

import "fmt"

func main() {
	// Type assertions for interface
	// https://go.dev/tour/methods/15
	// This is to verify that the value hold by a var of type interface (eg shape) is of a specific type (eg Circle1)
	var i interface{} = "hello" // i is of type *empty interface*, thus can hold a value of any type, and a string is assigned to it.
	// Now, if i holds a string, then assign it to a var s
	s := i.(string) // this would trigger a panic if i hold a value of different types other than string
	fmt.Println(s)

	// To avoid a panic, add a check var in the return of the assignment
	s1, isOk1 := i.(string) // if i holds a string, isOK is true, else false
	fmt.Println(s1, isOk1)  // hello true
	num, isOk2 := i.(int)   // i holds a string, not an int, so isOk2 is false, and num gets a default value 0
	fmt.Println(num, isOk2) // 0 false

	// if a boolean value is not present, a panic would trigger if type is wrong
	// num2 := i.(int) -->> this will trigger a panic
}
