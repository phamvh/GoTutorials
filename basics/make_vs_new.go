package main

import "fmt"

func main() {
	// https://www.godesignpatterns.com/2014/04/new-vs-make.html
	// new returns a pointer, make returns a value
	map1 := make(map[string]int, 4) // make allows an arg for capacity
	map2 := new(map[string]int)     // new does not allow capacity

	fmt.Println(map1) // map[]
	fmt.Println(map2) // &map[]

	// Note that make() can only be used to initialize slices, maps, and channels
	// make actually initializes a var to default value(s)
	slice1 := make([]int, 5)
	// new only allocates memory and returns a pointer without initialization
	slice2 := new([]int)
	fmt.Println(slice1) // [0 0 0 0 0]
	fmt.Println(slice2) // &[]
}
