package main

import "fmt"

func main1() {
	var b bool = true
	var num int = 10
	var f float32 = 3 / 14
	var s string = "hi"
	fmt.Println(b, num, f, s)

	var n1 int = 1
	var n2 int8 = 2
	var n3 int16 = -1
	var n4 int32 = 1
	var n5 int64 = -1
	var n6 uint = 1 // usigned integer
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
