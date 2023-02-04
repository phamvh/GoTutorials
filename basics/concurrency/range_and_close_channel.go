package main

import (
	"fmt"
	"time"
)

/**
https://go.dev/tour/concurrency/4
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch

ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.
Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the
receiver must be told there are no more values coming, such as to terminate a range loop.
*/

func fibonacci(n int, c chan int) {
	x, y := 1, 2
	for i := 0; i < n; i++ {
		time.Sleep(1 * time.Second)
		c <- x
		// both expressions on the right side (y, x+y) are evaluated first before their results
		// are assign to x and y ont he left side. This is the advantage of multiple assignments.
		x, y = y, x+y

	}
	close(c) // close the channel. Any attempt to send data into this channel after this will cause a panic.
}

func main() {
	c := make(chan int, 4)
	go fibonacci(cap(c), c)
	// Run the for loop in the main thread to make the main thread wait for the child thread
	for i := range c { // keep reading until the channel is closed
		fmt.Println(i)
	}

	// Another way to keep reading from the channel until it's closed
	c = make(chan int, 4)
	go fibonacci(cap(c), c)
	ok := true
	var v int
	for ok == true {
		// Note: I tried putting `v, ok = <-c` and it didn't work; it was not blocked and I ran into infinite loop
		// for v, ok = <- v; ok == true {... infinite loop here }
		// So one cannot put the receiving statement from a channel in the for condition, it seems.
		v, ok = <-c // ok is false if there are no more values to receive and the channel is closed.
		fmt.Println(v, ok)
	}
	/*
		v, ok := <-c
		fmt.Println(v, ok)
		v, ok = <-c
		fmt.Println(v, ok)
		v, ok = <-c
		fmt.Println(v, ok)
		v, ok = <-c
		fmt.Println(v, ok)
		v, ok = <-c
		fmt.Println(v, ok)

	*/
}
