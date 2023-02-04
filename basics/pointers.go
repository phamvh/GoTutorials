package main

import "fmt"

func main() {
	var num int = 10
	// Use ampersand & to access the address of a var
	// for info about %x, see: https://pkg.go.dev/fmt
	fmt.Printf("Address of a is %x\n", &num) //  c0000aa008

	// a pointer if a var, whose value is the address of another var, ie the memory location
	// Use star * to declare a var of pointer type
	// var var_name *var_type

	var ip *int
	var fp *float64
	i := 10
	f := 3.14
	ip = &i
	fp = &f
	// use * in front of a pointer to get the actual value stored at that address
	fmt.Println(i, f, ip, fp, *ip, *fp) // 10 3.14 0xc0000aa020 0xc0000aa028 10 3.14

	// nil pointer
	var b *int
	fmt.Println(b) // <nil>
	//fmt.Sprintf("the value of a nil pointer is %x\n", b) // -> this doesn't print anything. why???
	fmt.Println(b == nil) // true

	// array of pointers
	var nums = [3]int{100, 200, 300}
	var ptrs [3]*int
	for i := 0; i < len(nums); i++ {
		ptrs[i] = &nums[i]
	}
	fmt.Println(ptrs) // [0xc00001c090 0xc00001c098 0xc00001c0a0]
	for i := 0; i < len(ptrs); i++ {
		fmt.Print(*ptrs[i], ", ") // 100, 200, 300,
	}
	fmt.Println()
	for index := range ptrs {
		// range actually gives you the index of an array
		fmt.Print(index, ", ") // 0, 1, 2, -->> these are indices
	}
	fmt.Println()
	for index := range ptrs {
		fmt.Print(*ptrs[index], ", ") // 100, 200, 300,
	}
	fmt.Println()

	/////////////////////////////////
	// pointer to a pointer
	// A pointer to a pointer is a form of chain of pointers. Normally, a pointer contains the address of a variable.
	// When we define a pointer to a pointer, the first pointer contains the address of the second pointer,
	// which points to the location that contains the actual value

	var a int
	var ap *int
	// this is not de-pointering; in declaration it is a pointer to another pointer
	// However, in expression, like var c = *ap, it is actually de-pointering to get the value stored at ap.
	var app **int
	a = 1000
	ap = &a
	app = &ap
	// app = &&a  this (double ampersands) is not allowed :)
	fmt.Println(a)     // 100
	fmt.Println(*ap)   //  1000 de-pointering
	fmt.Println(**app) // 1000 double-de-pointering

	// passing pointers to a func
	num1, num2 := 1, 2
	fmt.Println(num1, num2) // 1 2
	swap(&num1, &num2)
	fmt.Println(num1, num2) // 2 1
}

func swap(a *int, b *int) {
	temp := *a // save the int value at address a
	// when a star is on the left side of the assignment, it means to put something into this address a.
	// Do NOT read *a on the left side as de-pointering. *b on the right side is de-pointering indeed.
	*a = *b   // no need to dig into the technicals of this; by Go's definition, it is putting the value at b into the address at a
	*b = temp // put value temp into the address at b
}
