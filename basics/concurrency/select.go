package main

import (
	"fmt"
	"time"
)

/*
*
In Go language, the select statement is just like switch statement, but in the select statement,
case statement refers to communication, i.e. sent or receive operation on the channel.

select{

case SendOrReceive1: // Statement
case SendOrReceive2: // Statement
case SendOrReceive3: // Statement
.......
default: // Statement

Select statement waits until the communication(send or receive operation) is prepared for some cases to begin.
The default case in a select is run if no other case is ready.
*/
func fibonacci1(c, quit chan int) { // two channels of type int
	x, y := 0, 1
	for { // for with no condition is infinite loop. Need break or return to get out
		// select statement lets goroutine wait on multiple communication operations
		// select blocks until one of the case runs, then it executes that case.
		// if multiple cases are ready at the same time, it chooses one randomly
		select {
		case c <- x: // send x to channel c
			x, y = y, x+y
		case <-quit: // read from channel quit, omit result?
			fmt.Println("quit")
			return
		}
		time.Sleep(time.Second)
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// run an anonymous func in a separate thread
	// Whenever you need to run a block of code in a thread, put code in an anonymous func like this.
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c) // try to read from channel c. This blocks until c has something.
			// this runs in separate thread, so t's OK. No deadlock.
		}
		// writing into quite channel make it avail for reading. Thus, in the fibo function,
		// the select statement has one more choice. Previously it can write into c only cuz quite has nothing to read from
		// now, it can either write into c, or read from quit. If it chooses the latter, func will return.
		quit <- 0 // write into quit channel
	}() // need these parens to call/run the anonymous func
	// run fibo function to either write into c, or quit the function.
	fibonacci1(c, quit)

	// another more understandable example of select
	// Here. we have 2 channels, and we write into these two channels after sleeping for 1 and 2 seconds.
	// Then we have the select with two cases trying to read from the channel. They block till any of them
	// is ready, and when one is ready, select will pick it and run it.
	c1, c2 := make(chan string), make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)

		// -->> if we have this default, this will execute immediately cuz the two writing goroutines above
		// sleep for 1 and 2 seconds, thus they are not ready yet. If this default is NOT present, select will block
		// until one of the cases above is ready. However, if this default is present, it gets executed immediately because
		// the two cases above are not ready.
		// default: fmt.Println("default") for a little
	}

}
