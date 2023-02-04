package main

import "fmt"

func main() {
	steps()
	fmt.Println()

	multiple_defers()

}

func steps() {
	// See here for details: https://gobyexample.com/defer
	step1()
	// Use defer keyword to defer the execution of a statement until the very last inside the enclosing function
	// Here, the enclosing function is steps, so even though step4_cleanup() is called after step1, before
	// step2 and step3, but it is executed after step2 and step 3.
	// The idea/goal of defer is to do cleanup.
	// Assume that in step 1 we open a file f.
	// Right after step 1, we want to make sure we don't forget to close the file f, so we call
	//     defer f.close()
	// in step 2 and 3, we can work with content of f

	// The main purpose of defer is to have it like a "finally" block in Java.
	// I intentionally made an error in step 2 by dividing a number by 0, and saw that
	// step4_cleanup still executed, but step 3 did not execute
	defer step4_cleanup()
	step2()
	step3()
}

func step1() {
	fmt.Print("step1 ")
}
func step2() {
	fmt.Print("step2 ")
	// intentionally make an error here
	b := 4 - 4
	a := 10 / b
	fmt.Println(a)
}
func step3() {
	fmt.Print("step3 ")
}

func step4_cleanup() {
	fmt.Print("step4 cleanup ")
}

func multiple_defers() {
	// when there are multiple defers, the defers are pushed into a stack, and get executed as first in last out
	for i := 1; i < 5; i++ {
		defer fmt.Print(i, ", ")
	}
	fmt.Println("Let's count from 5 to 1") // this will be called first, and then the deferred ones
	// the result of calling this func is:
	// Let's count from 5 to 1
	// 4, 3, 2, 1,
}
