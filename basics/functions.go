package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	log("Hello")
	fmt.Println(sum1(1, 2))

	theSum, theMessage := getSumAndMessage(1, 2, "John")
	fmt.Println(theSum, theMessage) // 3 Hello John

	var theSum1 int
	var theMessage1 string
	theSum1, theMessage1 = getSumAndMessage(1, 2, "Sam")
	fmt.Println(theSum1, theMessage1) // 3 Hello Sam

	// Ignore some value by using underscore
	mySum, _ := getSumAndMessage(2, 4, "OK")
	fmt.Println(mySum)

	fmt.Println(factorial(4))

	fmt.Println(squareNum(10)) // 100

	// pass a func as arg to another func
	fmt.Println(compute(1, 2, sum))     // 3
	fmt.Println(compute(1, 2, product)) // 2

}

func sum(a int, b int) int { // the last int is th e return type
	return a + b
}

func product(a, b int) int {
	return a * b
}

// pass a func as arg to another func
func compute(x int, y int, function func(int, int) int) int {
	return function(x, y)
}
func log(message string) {
	fmt.Println(message)
	return // this is optional when the function does not need to return anything
}

// named return value: give the return value a name, and then use the naked return
// without specifying the returned value
// Also note that when multiple cars share the same type, you can omit the type till the last one
func sum1(a, b int) (result int) { // (result in have to be in parens)
	result = a + b
	return // This is called a naked return. No need to put result here since go already knows what to return.
	// But you do need the "return" statement
	// of course, you can still put "return result" here and it is recommended to do so.
	// Do NOT use naked return often; only use it in a short function
}

// function can return multiple values.
// When calling this function, use similar syntax in Python to store multiple values
// theSum int, theMessage string = getSumAndMessage(1,2, "John")
func getSumAndMessage(a int, b int, name string) (result int, message string) {
	result = a + b
	message = "Hello " + name
	return
}

// The same as above, but we don't use named return vars here.
func getSumAdMessage2(a int, b int, name string) (int, string) {
	return a + b, "Hello " + name
}

// recursion. Go allows recursion
func factorial(n int) int {
	if n <= 1 {
		return n
	} else {
		return n * factorial(n-1)
	}

}

// a function syntax similar to javascript
func squareNum(num int) int {
	// function as a value, like in javascript. Because it is a statement, it needs to be inside another function
	getSquare := func(x int) int {
		return x * x
	}

	return getSquare(num)
}
