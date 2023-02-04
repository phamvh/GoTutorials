package main

import (
	"fmt"
	"math"
)

/*
*
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
Note that there is no class in Go, so you just need to implement this interface as a function

A type implements an interface by implementing its methods.
There is no explicit declaration of intent, no "implements" keyword.

See this for the under-the-hood: https://go.dev/tour/methods/11
*/

// Shape interface can define a set of methods. Here, just one method.
type Shape interface {
	area() float64 // any value that has this func can be thought of type Shape
}

// define a struct Circle
type Circle1 struct { // we already used name Circle in some other file of the same package, so use Circle1 here
	x, y, radius float64
}

// define a struct Rectangle
type Rectangle struct {
	width, height float64
}

// add a method area() to struct Cirle - not using interface yet
// Since circle now has method area(), it can be assigned to a var of type Shape.
// This method means type Circle1 implements the interface Shape,
// but we don't need to explicitly declare that it does so. https://go.dev/tour/methods/10
func (circle Circle1) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// add a method area() to struct Rectangle - not using interface yet
// we use pointer as a receiver here, not value, to add some spice.
func (rec *Rectangle) area() float64 {
	return rec.width * rec.height
}

// Now define a function getArea, which takes a shape as arg.
// This has nothing to do with enforcing Circle1 and Rectangle to be related to interface Shape
// However, we can pass a circle or a rectangle to this function, because they both have method area()
// So it is like a dynamic call, and it determines that the arg object actually has the method area()
// This is understandable, because there is no class in Go to implement this interface Shape, so any
// struct which has method area() is thought to implement Shape.
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle1{x: 0, y: 0, radius: 2.2}
	rec := Rectangle{2, 3.2} // implicit assignment of members. I guess it takes by the order
	// Even though, circle is not of type Shape here, it can still be passed to getArea. This is the idea of interface in Go.
	fmt.Println(getArea(circle)) // 15.205308443374602
	// Have to pass a pointer here as the receiver for Rectangle is a pointer
	fmt.Println(getArea(&rec)) // 6.4
	// However, you can call getArea() with a pointer to a circle.
	// so when the receiver is a value, you can use either value or pointer
	// but if the receiver is pointer, you can only use pointer.
	fmt.Println(getArea(&circle)) // 15.205308443374602

	// you can even assign these vars to a shape var
	var shape Shape
	shape = circle
	fmt.Println(shape.area())
	shape = &rec // also pointer
	fmt.Println(shape.area())

	// Note that the receiver of Circle method is a value, not a pointer.
	// So you have to pass/assign a value of Circle, not a pointer to a Circle to shape.
	shape = &circle
	shape.area()

}
