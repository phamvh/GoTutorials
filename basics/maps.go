package main

import (
	"fmt"
)

func main() {
	/**
	There are two ways to declare a map: the type of key is in brackets, indicating that it is used as index to find value
	var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	*/
	var a = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(a)          // map[1:one 2:two 3:three]
	fmt.Println(a[1], a[2]) // one two

	// create a map and init to empty
	var b = map[int]string{}
	b[1] = "one"
	b[2] = "two"
	fmt.Println(b) // map[1:one 2:two]

	// *Declare* a nil (null) map without using braces. To do this, do NOT use = operator
	var b2 map[int]string // see? no = used here. Weird though
	//b2[10] = "ten" -> not allowed as the map is null now.
	fmt.Println(b2 == nil) // true

	// To create a map that is not nil, you can either use the syntax in line 18 to init it to empty {}
	// or use the make function as follows:
	// see here for differences when creating map using and not using make():
	// https://stackoverflow.com/questions/16959992/creating-map-with-without-make
	// Technically, make allows you to specify the size, while the other way doesn't.
	// 1. So when you create a map with initial values, you must use  var b = map[int]string{}
	// 2. When creating an empty map, use either way as they are equivalent
	// 3. When creating a map with size, have to use make()
	c := make(map[int]string) // no need to provide {} here, and this is not nil. You can assign values to it.
	c[3] = "three"
	c[4] = "four"
	fmt.Println(c) // map[3:three 4:four]

	// update elements in a map
	var d = map[int]string{1: "one", 2: "two", 3: "three"}
	d[1] = "ONE"
	// remove an element
	delete(d, 3)
	fmt.Println(d) // map[1:ONE 2:two]

	///////////////////////////////////////
	// Check if a a key is in a map or not
	var e = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	val1, isPresent1 := e["one"]
	fmt.Println(val1, isPresent1) // 1 true
	val5, isPresent5 := e["five"]
	fmt.Println(val5, isPresent5) // 0 false // note that 0 is returned cuz it's a default value for int, even though key does not exist
	// omit value
	_, isPresent4 := e["four"]
	fmt.Println(isPresent4) // true

	// iterating over a map. Hint: use range keyword
	e = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for key, val := range e {
		fmt.Print(key, ":", val, ", ")
	}
	fmt.Println()

	// iterate over keys only
	for key := range e {
		fmt.Print(key, ",") // two,three,four,one,
	}
	fmt.Println()
	// there is not a short function to get all keys from a map.
	// one needs to create a slice first, and then assign each key to it
	i := 0
	keySlice := make([]string, len(e))
	for key := range e {
		keySlice[i] = key
		i++
	}
	fmt.Println(keySlice) // [one two three four]

	// get all values in a similar way
	// WRONG result. See below for why, and correction
	values := make([]int, len(e), len(e))
	for _, val := range e {
		values = append(values, val)
	}
	fmt.Println(values) // [0 0 0 0 3 4 1 2] See something WRONG????
	// THis is because in make([]int, len(e), len(e)), we set the actual length to len(e) already, and therefore
	// the append() function will start appending from position len(e) instead from 0.
	// TO make this correct, we need to specify the len to be 0, and cap to b len(e).
	// Note that if cap is not sufficient, new underlying array will be allocated to support the map, so pay attention
	// to efficiency here.
	values = make([]int, 0, len(e))
	for _, val := range e {
		values = append(values, val)
	}
	fmt.Println(values) // [1 2 3 4] -> correct as expected

}
