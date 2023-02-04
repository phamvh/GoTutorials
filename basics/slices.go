package main

import "fmt"

/*
*
Slices are like list in java and python. They are dynamic: size may shrink or grow.
Can only store elements of the same type.
A slice is internally just a pointer to the backing array. It is a reference type.
The difference in syntax between an array and a slice is:
  - array: need to indicate the size, or use three dots: var arr = [3]int{1,2,3}, or var arr = [..]int{1,2,3}
  - slice: should not indicate the length in the square brackets: var sli = []int{1,2,3}

The basic difference between a slice and an array is that a slice is a reference to a contiguous segment of an array.
A slice is ALWAYS associated with a backing array.
Unlike an array, which is a value-type, slice is a reference type. A slice can be a complete array or a part of an array,
indicated by the start and end index. A slice, therefore, is also an array that pours a context of dynamism onto the
underlying array, which otherwise is a static contiguous memory allocation.
https://www.developer.com/languages/arrays-slices-golang/

According to this: https://go.dev/tour/moretypes/8
1. slice is much more common than array in Go
2. A slice does not store any data, it just describes a section of an underlying array.
3. Changing the elements of a slice modifies the corresponding elements of its underlying array.
4. Other slices that share the same underlying array will see those changes.
*/
func main() {
	// compared to array's declaration, you don't give the size for slices here.
	/**
	A slice can be declared/created as follows:
	make([]Type, length, capacity)
	make([]Type, length)
	[]Type{}
	[]Type{value1, value2, ..., valueN}
	*/
	s1 := []int{} // init an empty slice
	// this is actually 2 steps: create an array, and make the slice reference the entire array.
	s2 := []int{1, 2, 3} // should NOT indicate the length in [], otherwise it would be perceived by GO as an array instead.

	fmt.Println(len(s1), len(s2)) // 0, 3 : actual size
	fmt.Println(cap(s1), cap(s2)) // 0, 3 : capacity - the number of elements the slice can grow or shrink to

	arr := [5]int{1, 2, 3, 4, 5}      // this is an array, because we put length inside []
	s3 := arr[2:4]                    // [3 4]: create a slice from an array [i:j]: including i, excluding j
	fmt.Println(s3, len(s3), cap(s3)) // [3 4] 2 3: the slice can grow to the end of the array, thus cap = 3 only as it started from 2nd elem

	// use make() function to create a slice. this is how you create dynamically-sized arrays.
	// this creates an array of size 10, and make the slice s4 reference to it.
	s4 := make([]int, 5, 10)          // actual size = 5, capacity = 10
	fmt.Println(s4, len(s4), cap(s4)) // [0 0 0 0 0] 5 10
	s5 := make([]int, 5)              // omit capacity, so now cap = len
	fmt.Println(s5, len(s5), cap(s5)) // [0 0 0 0 0] 5 5

	// access elements in a slice: index starts from 0
	s6 := []int{1, 2, 3, 4}
	s6[3] = 20
	fmt.Println(s6[0], s6[1], s6[3]) // 1 2 20

	// append elements to a slice
	// Note that this does NOT modify the original slice s6. It actually creates a new slice, and appends elements to it.
	s6 = []int{1, 2, 3, 4}
	s6_appended := append(s6, 5, 6)
	fmt.Println("s6 after appended", s6)    // s6 after appended [1 2 3 4] -->> See? not changed
	fmt.Println("s6_appended", s6_appended) // s6_appended [1 2 3 4 5 6]

	// append two slices
	s7 := []int{7, 8}
	s8 := append(s6, s7...) // the three post dots are like unpacking
	fmt.Println(s8)         // [1 2 3 4 5 6 7 8]

	// you can append more elements than the capacity of a slice. It will grow to accommodate.
	arr = [5]int{1, 2, 3, 4, 5}
	s9 := arr[3:4]
	fmt.Println(s9, len(s9), cap(s9)) // [4] 1 2, so cap is 2.
	// let's append 4 more elements
	s9 = append(s9, 5, 6, 7, 8)
	fmt.Println(s9, len(s9), cap(s9)) // [4 5 6 7 8] 5 6 // cap now is 6

	//  MEMORY EFFICIENCY
	// Go loads all numbers of an array or slice into memory. When the size is large, one can use
	// copy() function to only load the elements needed for work.
	// The copy(dest, src) function takes in two slices dest and src, and copies data from src to dest.
	//It returns the number of elements copied.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	neededNumbers := numbers[:len(numbers)-10]                         // slice from the beginning (omitting the first index) till j
	fmt.Println(neededNumbers, len(neededNumbers), cap(neededNumbers)) // [1 2 3 4 5] 5 15

	copy1 := make([]int, len(neededNumbers)) // create a slice of type int given the length, NOT given capacity
	copiedLength := copy(copy1, neededNumbers)
	fmt.Println(copy1, len(copy1), cap(copy1)) // [1 2 3 4 5] 5 5
	fmt.Println(copiedLength)                  // 5

	copy2 := make([]int, 10)
	copy(copy2, numbers)                       // copy from the beginning of src (numbers) till all the dest (copy2) get filled
	fmt.Println(copy2, len(copy2), cap(copy2)) // [1 2 3 4 5 6 7 8 9 10] 10 10

}
