package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hi Huy! Welcome to GoLang")

	// constants
	const number int = 1
	const st string = "hello"

	// variables
	var n int = 10
	var m = 20 // type inferred
	// assign operator can only be used within a func
	hello := "Hello" // use assign operator, "var" is NOT allowed in front. Can only be used inside a func
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(hello)

	// default values if no initial values given
	var str string
	var num int
	var b bool
	fmt.Println(str) // empty string
	fmt.Println(num) // 0
	fmt.Println(b)   // false

	// multiple vars without init.
	//Format: var variable_list optional_data_type;
	var s1, s2 string      // can only be used with vars of same type. Can't do: "var i int, s string"
	s1, s2 = "hi", "hello" // assign multiple values to multiple vars. Don't do this though. looks confusing
	// just do: s1 = "hi", s2 = "hello"
	fmt.Println(s1, s2) //hi hello

	// multiple vars with init
	var num1, str1 = 10, "Hello there"
	num2, str2 := 10, "Hello baby" // Note that "var" is not allowed before the assign operator
	fmt.Println(num1+num2, str1+str2)

	// block declaration
	var (
		num3 int
		num4 int    = 10
		str3 string = "hello"
	)
	fmt.Println(num3, num4, str3)

	// const; note that const cannot be declared with :=, only with =
	const PI = 3.14
	const EXP float32 = 2.78
	fmt.Println(PI, EXP)

	// get the type of a variable
	// There are 3 ways
	x := 3.14

	// use refect package
	fmt.Println(reflect.TypeOf(x)) // float64

	// Use string format
	var theType string = fmt.Sprintf("%T", x)
	fmt.Println(theType) //  float64

	// Use switch for type of interface only
	/*
		switch x.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("unknown type")
		}
	*/

}
