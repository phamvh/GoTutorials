package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// To run a function-call in a separate thread, put go in front of it.
	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	//The sync package provides useful primitives, although you won't need them much in Go as there are
	// other primitives.
	go say("Hello")

	// We do NOT run this in a separate thread. Rather, we run this in the main thread
	// so make the main thread wait  for the goroutine above to finish.
	say("World")

}
