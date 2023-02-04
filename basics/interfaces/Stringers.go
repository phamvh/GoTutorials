package main

import "fmt"

/*
https://go.dev/tour/methods/17
This is an interface that is similar to toString() in Java.
It defines one method, String(), that returns a string (supposedly representing a value)

	type Stringer interface {
	    String() string
	}
*/

// Person /**
type Person struct {
	Name string
	Age  int
}

// now add method String to Person
func (p Person) String() string {
	// note that Sprintf formats and returns a string. It does not print it like Printf.
	return fmt.Sprintf("%v : %v years", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// Now, Println automatically calls the String() method of Person, just like in Java calling toString()
	fmt.Println(a) // Arthur Dent : 42 years
	fmt.Println(z) // Zaphod Beeblebrox : 9001 years
}
