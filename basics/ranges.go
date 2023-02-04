package main

import "fmt"

func main() {
	/**
	1. range is a keyword, not a function
	2. It is used to iterate over items of an array, slice, channel or a map
	3. with arrays, it returns the index of the item as integer. You can also return
	   both index and val while iterating a slice.
	4. with map, it returns key map key of the next key-value pair
	*/
	numbers := []int{1, 2, 3, 4, 5} // create a slice (Note: this is NOT an array)
	for i := range numbers {
		fmt.Print(numbers[i], ", ") // 1, 2, 3, 4, 5,
	}
	fmt.Println()

	// getting both index and value from range
	for i, v := range numbers {
		fmt.Printf("%d : %v, ", i, v) // 0 : 1, 1 : 2, 2 : 3, 3 : 4, 4 : 5,
	}
	fmt.Println()

	mymap := map[int]string{1: "one", 2: "two"}
	for key := range mymap {
		fmt.Print(key, ":", mymap[key], ", ") // 1:one, 2:two,
	}
	fmt.Println()
	// range can actually return both key and value like in python
	// you just need to indicate it like below:
	// use `for _,val := range ...` to ignore key
	for key, val := range mymap {
		fmt.Print(key, ":", val, ", ") // 1:one, 2:two,
	}
	fmt.Println()
}
