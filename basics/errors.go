package main

import (
	"errors"
	"fmt"
	"math"
)

/**
Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration −

type error interface {
   Error() string
}
*/

// When a function may encounter an error, it can choose to return multiple values
// where the error value is usually the last item in the returned values.
func Sqrt(num float64) (float64, error) {
	// the last returned value is of type error. This means, any struct that has the Error method.
	if num < 0 {
		// use package errors to create a new error object
		return 0, errors.New("Num cannot be less than 0")
	} else {
		return math.Sqrt(num), nil // nil here is for error
	}
}

func main() {
	val, err := Sqrt(-1)
	if err == nil {
		fmt.Println("square root is ", val)
	} else {
		fmt.Println("Error encountered", err) // Error encountered Num cannot be less than 0
	}
}
