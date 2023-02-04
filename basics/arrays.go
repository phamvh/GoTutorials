package main

import "fmt"

func main() {
	/**
	An arrray can be declared using the following syntax:
	[N]Type
	[N]Type{value1, value2, ..., valueN}
	[...]Type{value1, value2, ..., valueN}
	Note that we need to either provide size N, or use three dots ... to ask Go to infer the length.
	Do NOT confuse this with slice, which does NOT have size in [].
	If you write `[]TYPE`, then it's a slice, not an array.
	A slice is a reference to an array in essence.
	*/
	// declare an array of type int, and length = 10
	// it can be: `a0 [10] int`, `a0[10] int` or `a0 [10]int1
	var a0 [10]int
	a0[0] = 1
	fmt.Println(a0) // [1 0 0 0 0 0 0 0 0 0]

	var a1 = [3]int{1, 2, 3} // size 3, and init to 3 elements
	a2 := [3]int{4, 5, 6}    // use assign, so no "var"

	fmt.Println(a1, a2)

	// inferred length
	var a3 = [...]int{1, 2, 3} // inferred size
	a4 := [...]int{4, 5, 6}
	var a5 = [2]string{"hi", "bye"}
	fmt.Println(a3, a4, a5)
	// var error_arr = [2]int{1, 2, 3} //  index 2 is out of bounds

	fmt.Println(a3[0], a4[1], a5[0])
	a5[1] = "hey"
	fmt.Println(a5)

	a6 := [5]int{}                     // not initialized, all defaulted to 0
	a7 := [5]int{1, 2}                 // partially inited, other defaulted to 0
	a8 := [5]int{1, 2, 3, 4, 5}        // fully inited
	a9 := [5]string{0: "Hi", 2: "Hey"} // inited specific elements
	fmt.Println(a6, a7, a8)            // [0 0 0 0 0] [1 2 0 0 0] [1 2 3 4 5]
	fmt.Println(a9)                    // [Hi  Hey  ]

	fmt.Println(len(a9)) // 5

	// multi-D arrays
	var a10 [2][3]int // 2-D array of type int, 2 rows and 3 cols
	a10[0][0] = 1
	fmt.Println(a10) // [[1 0 0] [0 0 0]]

	a11 := [2][2]int{
		{1, 2},
		{2, 3},
	}
	fmt.Println(a11) // [[1 2] [2 3]]

	// Updating an array when passing it to a function
	// There are 3 ways:
	//  passing the array, and pointer to the array, or an array of pointers, each points to an element in the array
	// All three ways can update the array.
	arrayForUpdate := []int{1, 2, 3}
	fmt.Println(arrayForUpdate) // [1 2 3]
	// passing the array to a func
	updateArray(arrayForUpdate)
	fmt.Println(arrayForUpdate) // [1000 2 3]

	arrayForUpdate = []int{1, 2, 3}
	fmt.Println(arrayForUpdate) // [1 2 3]
	// passing a pointer to the array to a func
	updateArrayPointer(&arrayForUpdate)
	fmt.Println(arrayForUpdate) // [1000 2 3]

	arrayForUpdate = []int{1, 2, 3}
	// Note, I need to use make() to make an SLICE here.
	// I first did `var arrayOfPointers []*int`, then this returns an array of len = 0, and got index out of bound.
	var arrayOfPointers = make([]*int, len(arrayForUpdate), len(arrayForUpdate))
	for i := 0; i < len(arrayForUpdate); i++ {
		arrayOfPointers[i] = &(arrayForUpdate[i])
	}
	fmt.Println(arrayForUpdate) // [1 2 3]
	// passing an array of pointers, each of which points to an element in the original array
	updateArrayOfPointers(arrayOfPointers)
	fmt.Println(arrayForUpdate) // [1000 2 3]

}

// passing an array
/**
NOte that the arg is []int - an array of int of UNKNOWN length, so when you call this func, you also need to
pass an array that was declared WITHOUT length, otherwise you would get incompatible arg.
For example:
  - this is NOT ok: var arr = [3]int{1,2,3}
                    updateArray(arr) -->> this causes an error  because [3]int is not compatible with []int.
After reading some on internet, I believe that `a []int` indicates a SLICE of int, not an array, while
[3]int{1,2,3} is actually an array, not a slice. So that's the reason for incompatibility.
*/
func updateArray(a []int) {
	a[0] = 1000 // this does update the array after func exits
}

// passing a pointer to an array
func updateArrayPointer(pa *[]int) {
	// first, (*pa) gives the array, then use the index to update an element
	(*pa)[0] = 1000 // // this does update the array after func exits
}

// passing an array of pointers
func updateArrayOfPointers(ptrs []*int) {
	// *ptrs[0] gives you a pointer, so need to a star in the front to assign a value to it.
	*(ptrs[0]) = 1000 // // this does update the array after func exits
}
