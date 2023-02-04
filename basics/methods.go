package main

import (
	"fmt"
	"math"
)

// define a struct for a Circle
type Circle struct {
	x, y, radius float64
}

/*
*
Now we want to add a method to the struct above.
The syntax looks a little unclear though.
It is like:
  - an anonymous function that takes a parameter named circle of type Circle
  - the anonymous function returns another function, named area, which in turn returns a float64
    This is s bit similar to a closure, which returns an anonymous function where the return type is "func() float64", not "are() float64"
  - each object of type Circle now will have a method area(), in addition to x, y and radius members.
*/
// circle here is called the Receiver arg.
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/**
You can define a method where the receiver arg is a pointer too, not just the struct itself.
https://go.dev/tour/methods/4
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Note that when a struct is passed to a func, Go creates a copy of it and passes the copy to the func.
So no modifications can be made on the original struct obj when passing it to a func.
However, when a pointer of a struct is passed to a func, updates can be made to the struct obj.

One cool thing is, you can call this method with either a value or a pointer of Circle.
Go interprets circle.shrinkCircleByHalf() as (&circle).shrinkCircleByHalf().
In general, a method with a value receiver or pointer receiver, can be called from a value or a pointer, and vice versa.
It's just a flexibility that Go provides.

As stated here: https://go.dev/tour/methods/8
One should use pointer receiver when creating a method to avoid copy of value each time a method is called.
In addition, a pointer receiver can modify the original value.
*/

func (circlePtr *Circle) shrinkCircleByHalf() {
	// because circlePtr is a pointer, this udpates the radius even after this method exits.
	circlePtr.radius = circlePtr.radius / 2
}

// a more general version of shrink where you can pass more args.
// all additional args must be passed through the second pair of parens, not the first pair.
func (circlePtr *Circle) shrink(factor int /*addition args are passed here*/) {
	circlePtr.radius = circlePtr.radius / float64(factor)
}

// You can declare a method on non-struct type too.
// https://go.dev/tour/methods/3
// Let's define our own type
type MyFloat float64

// let's add a method to this type
// Note: You can only declare a method with a receiver whose type is defined in the same
// package as the method. You cannot declare a method with a receiver whose type is defined
// in another package (which includes the built-in types such as int).
func (myFloat MyFloat) abs() float64 {
	if myFloat < 0 {
		return float64(-myFloat)
	}
	return float64(myFloat)
}

func main() {
	circle := Circle{
		x:      0,
		y:      0,
		radius: 5,
	}
	fmt.Println(circle.area()) // call method area of object circle. -> 78.53981633974483

	circlePtr := &circle
	circlePtr.shrinkCircleByHalf()
	fmt.Println(circle) // {0 0 2.5} radius is reduced by half.
	// Go also allows you to call shrinkCircleByHalf() bby value, not by pointer
	// so the following is fine:
	circle.shrinkCircleByHalf() // Note that you CANNOT do this with regular functions though. Only Ok for methods.
	fmt.Println(circle)         // {0 0 1.25}
}
