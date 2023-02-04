package main

import "fmt"

func main() {
	/**
	Use type_name(expression) to cast a value or expression to a type.
	*/
	var num1 int = 19
	var num2 int = 5
	var result float64
	// result = num1 / num2 -->> error: use num1 / num2 (value of type int) as type float64 in assignment
	// result = float64(num1) / num2 -->> still error: invalid operation: float64(num1) / num2 (mismatched types float64 and int)
	// Need to cast both num1 and num2 to float64
	result = float64(num1) / float64(num2)
	fmt.Println(result) // 3.8
}
