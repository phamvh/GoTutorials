package main

import "fmt"

/**
A closure is the combination of a function bundled together (enclosed) with references to its surrounding state.
A function closure is a function that returns another function.

In other words: https://go.dev/tour/moretypes/25
Go functions may be closures. A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

The purpose of function closures is they allow you to attach variables to an execution context.
Variables in closures can help you maintain a state that you can use later. They provide data encapsulation.
They help remove redundant code.
Each time the inner func gets executed, it still remembers the prev state stored in those vars.
*/

// note the return type "func() int", indicating it returns an anonymous function which returns int.
// Note that it can also return a named function, which is called a method. See methods for detials.
func getSequence() func() int {
	start := 0 // we want to hide this start var
	return func() int { // this is an inner and anonymous func
		start++
		return start // each time this inner func runs, it remembers the last value of start. This is a *closure*
	}
}

func main() {
	myFunc := getSequence()
	fmt.Println(myFunc()) // 1
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
}
