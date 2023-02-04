package main

import "fmt"

// Person struct is used to group data of different types together
type Person struct {
	name   string
	age    int
	salary int
}

func main() {
	var p1 Person
	p1.name = "John"
	p1.salary = 100000
	p1.age = 40

	fmt.Println(p1) // {John 40 100000}
	fmt.Println(p1.name, p1.age, p1.salary)
	printPerson(p1)

	// variable_name := structure_variable_type {value1, value2...valuen}
	var p2 Person
	p2 = Person{
		name:   "Mike",
		age:    33,
		salary: 100000, // this last comma is required if each member is written in a separate line.
	}
	fmt.Println(p2) // {Mike 33 100000}

	// Create a struct value with all members initiated to default values
	var defaultP Person                      // no equal sign, no braces
	fmt.Println("default struct:", defaultP) // default struct: { 0 0} -->> strings are empty, numbers are 0 by default
	// The above is equivalent to
	defaultP2 := Person{}  // give empty braces
	fmt.Println(defaultP2) // { 0 0}
	// you can compare two struct values. Go compares their members respectively
	fmt.Println(defaultP == defaultP2) // true

	// when creating a map of type struct, you can omit the struct name
	var peopleMap = map[int]Person{
		1:/* we do NOT need to put Person before the braces here */ {"John", 20, 10000},
		2: {"Lance", 22, 200000},
	}
	fmt.Println(peopleMap) // map[1:{John 20 10000} 2:{Lance 22 200000}]

	// pointer to a structure
	var pptr *Person
	pptr = &p2
	fmt.Println(*pptr) // {Mike 33 100000}
	// interestingly enough, use pointer in the same way as the struct var to access data inside a struct
	// so, pptr.name is the same as p2.name. In C++ you would need pptr->name or something like that.
	//
	// Just found out that, technically one need to do (*pptr).name to get name, but that is cumbersome
	// so Go allows you to use dot, pptr.name, without explicitly de-referencing a pointer.
	// https://go.dev/tour/moretypes/4
	fmt.Println(pptr.name, pptr.age, pptr.salary) // Mike 33 100000

	// or simply using the assign operator
	pptr2 := &p2
	fmt.Println(*pptr2) // {Mike 33 100000}

	// change contents of a struct when passing it to a func
	fmt.Println("Before passing to a func", p2) // Before passing to a func {Mike 33 100000}
	// struct is passed by value!!!!
	updatePerson(p2)                           // this will have NO effect on p2
	fmt.Println("After passing to a func", p2) // After passing to a func {Mike 33 100000} -->> did NOT update

	// now passing a pointer to p2 to a func
	fmt.Println("Before passing pointer to a func", p2) // Before passing pointer to a func {Mike 33 100000}
	updatePersonPointer(&p2)                            // this will have EFFECT on p2
	fmt.Println("After passing pointer to a func", p2)  // After passing pointer to a func {Alex 33 100000}
}

// passing a var of type Person will NOT change the value of the var
// struct is passed by value!!!!
func updatePerson(p Person) {
	p.name = "Alex"
}

// passing a pointer to a person will update the value of the var that the pointer points to
func updatePersonPointer(p *Person) {
	p.name = "Alex"
}

// pass a struct as arg
func printPerson(p Person) {
	fmt.Println(p.name, p.age, p.salary)
}
