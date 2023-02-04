package main

import (
	"fmt"
	"io"
	"strings"
)

// https://go.dev/tour/methods/21
// Simply copied from there.
// This example shows that the io.Reader interface has Read method, which returns both the number being read, and error if
// EOF is reached
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // for a slice of bytes, %q gives a double-quoted string safely escaped with Go syntax
		if err == io.EOF {
			break
		}
	}

	// Testing the MyReader below, which emits an infinite stream of 'A' chars
	myReader := MyReader{}
	i := 0
	sli := make([]byte, 8)
	for i < 5 {
		read, err := myReader.Read(sli)
		fmt.Printf("%q\n", sli[:read])
		if err != nil {
			break
		}
		i++

	}

}

// https://go.dev/tour/methods/22
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(sli []byte) (int, error) {
	var b = byte('A')
	for i := 0; i < len(sli); i++ {
		sli[i] = b
	}
	return len(sli), nil
}
