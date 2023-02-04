package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
String is a slice in Go -  read-only slice of bytes.
See this doc for details about Strings, bytes, runes and characters in Go:
   https://go.dev/blog/strings
*/

func main() {
	var str = "Hello World!"
	fmt.Printf("%s", str)
	fmt.Println()
	fmt.Println(len(str)) // 12

	// concatenate two strings
	var str2 = "How are you?"
	fmt.Println(str + str2) // Hello World!How are you?

	///////////////////////////////////////////
	// strings package has a vast amount of functions to work on strings
	// practically all methods for string in Java

	fmt.Println(strings.Compare(str, str2))           //  -1 (value can be -1, 0, 1
	fmt.Println(strings.Contains(str, "World"))       // true
	fmt.Println(strings.ContainsAny(str, "bahamama")) // false
	// check if str contains any character in the second str
	fmt.Println(strings.ContainsAny(str, "oh my god")) // true
	fmt.Println(strings.Index(str, "World"))           // 6
	fmt.Println(strings.Split(str, " "))               // [Hello World!]

	strs := []string{"hi", "there", "how"}
	fmt.Println(strings.Join(strs, "-")) // hi-there-how

	/////////////////////////////////
	// Rune in go. See this for details:
	//     https://dev.to/robogeek95/how-to-access-string-values-in-go-16cb
	// In Go, string is a slice of bytes. However, one character may occupy more than one byte.
	// Therefore, you CANNOT assume that each byte is one char.
	// Example
	testString := "Señor"
	fmt.Println(len(testString)) // 6   See???? It's not 4, but 6. It's 6 bytes, not chars. That's why a string a slice of bytes, not of chars
	for i := 0; i < len(testString); i++ {
		// %c means to print a byte as a char
		fmt.Printf("%c  ", testString[i]) // S  e  Ã  ±  o  r
	}
	fmt.Println()
	for i := 0; i < len(testString); i++ {
		// %x means to print a byte as a byte
		fmt.Printf("%x  ", testString[i]) // 53  65  c3  b1  6f  72
	}
	fmt.Println()
	// So, how to print chars then?
	// In Go, a rune is used to represent a char. A rune is just an int32; it's named rune to avoid confusion with the general int32.
	// A rune can occupy more than one byte
	// create an array of runes from the string
	runes := []rune(testString) // just pass the string with parens
	for i := 0; i < len(runes); i++ {
		// %c means to print a byte as a char
		fmt.Printf("%c  ", runes[i]) // S  e  ñ  o  r  -->> CORRECT
	}
	fmt.Println()
	// another way is to use the builtin function string to convert a rune to a char
	// Again, do NOT try to convert a single byte to a char, because a char, represent by a rune, can occupy more than 1 byte.
	for i := 0; i < len(runes); i++ {
		// no formatting needed, simply use Print, not Printf
		fmt.Print(string(runes[i]), " ") // S  e  ñ  o  r  -->> CORRECT
	}
	fmt.Println()

	/////////////////////////////////
	// Package unicode provides functions to work with chars, such as isLower, isUpper
	char := 'c'                         // use single quotes for a rune (char)
	fmt.Println(unicode.IsLower(char))  // true
	fmt.Println(unicode.IsDigit(char))  // false
	fmt.Println(unicode.IsLetter(char)) // true
	fmt.Println(unicode.IsSpace(char))  // false
}
