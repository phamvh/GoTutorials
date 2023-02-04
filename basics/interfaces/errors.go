package main

import (
	"fmt"
	"math"
	"time"
)

/**
https://go.dev/tour/methods/19
The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}

As with fmt.Stringer, the fmt package looks for the error interface when printing values.

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.
*/

type MyError struct {
	When time.Time
	What string
}

// make MyError implement interface *error* by adding method Error() to it
func (e *MyError) Error() string {
	return fmt.Sprintf("At %v, %s", e.When, e.What)
}

// just make a func that returns a value of type error
func run() error {
	return &MyError{
		time.Now(),
		"Some error occured",
	}
}

// A more real example
// https://go.dev/tour/methods/20

// ErrNegativeSqrt Create a custom type for float64.
// This type should implement interface error to return an error message.
// Later, we can create a value of this type to represent an error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Number %f is negative. It must be non-negative", float64(e))
}

// just experiment: make ErrNegativeSqrt implement Stringer also, to see which one fmt.Println picks: String() or Error()
// You should not do this though, cuz making ErrNegativeSqrt implement both error and Stringer interfaces is confusing.
// So the result of the experiment is, fmt.Println calls Error(), not String(), which makes sense because error has precedence.
func (e ErrNegativeSqrt) String() string {
	return "Oh no - negative number"
}

// Sqrt now create a function Sqrt
// if x is negative, return a non-nil error.
// Since our ErrNegativeSqrt implements interface error, we can use it in the place of the error here
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) //  a non-nil error is returned
	} else {
		return math.Sqrt(x), nil
	}

}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("It worked")
	}

	// the real example above
	result, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err) // Println looks for method Error() of value err to call to get the error message here
	} else {
		fmt.Println("result is ", result)
	}
}
