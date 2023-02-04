package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println()
	for i := 0; i < 50; i += 10 {
		if i == 20 {
			continue
		}
		if i == 40 {
			break
		}
		fmt.Print(i, ",") // 0,10,30,
	}
	fmt.Println()

	// the init and post statements in for loop are optional
	sum := 2
	for sum < 10 {
		sum = sum * 2
		fmt.Print(sum, " ") // 4 8 16
	}
	fmt.Println()

	/// infinite loop
	// for {
	// }
	// or
	// for ;; {}

	// The range keyword is used to more easily iterate over an array, slice or map.
	//It returns both the index and the value.
	// for index, value := array|slice|map {...}
	fruits := []string{"apple", "orange", "kiwi"}
	for idx, val := range fruits {
		fmt.Print(idx, ":", val, "; ") // 0:apple; 1:orange; 2:kiwi;
	}
	fmt.Println()
	// omit the index by using underscore, like in python
	for _, val := range fruits {
		fmt.Print(val, ", ") // apple, orange, kiwi,
	}

}
