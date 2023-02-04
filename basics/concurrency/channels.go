package main

import (
	"fmt"
	"time"
)

/**
A channel is a tube/conduit through which data can be sent and received.
Channel acts like a shared resource for two threads (goroutines) to communicate with each other.
A channel of type T only allows data of type T to be sent and received.
     to send to a channel ch: ch <- v
     to receive from ch     : v := <- ch
The data flows in the direction of the arrow

To make a channel of type int, use make() func
     ch := make(chan int) // chan is a keyword in GO for channel

Channel is used in thread programming.
*/

/**
By default, sends and receives block until the other side is ready. This allows goroutines to
synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines.
Once both goroutines have completed their computation, it calculates the final result.
*/

// sum compute the sum of elements in a slice and send it to a channel
func sum(s []int, c chan int) {
	res := 0
	for _, val := range s {
		res += val
	}
	time.Sleep(2 * time.Second)
	c <- res // send result to channel c
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int) // create a channel of type int so int can be sent into it, and received from it
	// compute sum for each half of the slice in different threads
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// now receive values from the channels. THis call is blocked until the data is ready in the channel
	// In other words, when data is not ready, a call to receive data from the channel waits till data is avai;
	sum1 := <-c // get the first data
	sum2 := <-c // get the second data; we know because we sent two pieces of data (int)
	fmt.Println(sum1 + sum2)

	// Buffered channel
	// A buffered channel is a channel with size. WHen the channel is full sending data to it is blocked.
	// Use make and provide a second arg as size to create a buffered channel
	buf := make(chan int, 2)
	// let's send 3 numbers into it
	buf <- 1
	buf <- 2
	// this call causes a deadlock
	buf <- 3 // this get blocked here, and waits forever here because we do not use threads when putting data in

	fmt.Println(<-buf)

}
