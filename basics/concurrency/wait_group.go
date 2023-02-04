package main

import (
	"fmt"
	"sync"
)

/**
https://www.geeksforgeeks.org/using-waitgroup-in-golang/
a WaitGroup is exactly a countdown latch.
1. It has an internal counter keeping track of how many works/threads/routines are currently deployed.
   Add(int): add the number of works
2. It has method Done(), which decrements the internal counter by 1: meaning one work is done.
3. Wait() method blocks the execution until the internal counter is 0.
*/

// Example  below is taken from here: https://golangbot.com/mutex/
// The main idea is, we have one var x, and 1000 goroutines that try to increment it.
// We want to achieve two things:
//   1. No two goroutines can increment the var at the same time. Use Lock/Unlock or a channel for this.
//   2. Make the main thread wait until all 1000 child threads finish. Use WaitGroup for this.

var x = 0 // this is the var we want to handle racing condition on.

// Increment function increases the value of x by 1.
// It is called by goroutines to run in parallel
// Note that it is important to pass the mutex as a pointer here. Otherwise, if we pass it as
// a value, each goroutine will have its own copy of the mutex and race condition still occurs.
// Similar for WaitGroup, as both WaitGroup and Mutex are struct.
// Note: I tried passing mutex as value, and got different result each time due to race condition.
//
//	I also tried passing waitGroup as value, and got a deadlock. So each time Go creates a copy
//	of waitGroup (if passed by value), therefore the Done() is called on the copy, not the original
//	waitGroup, so the original waitGroup waits forever as its counter never decreases.
func Increment(x *int, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	// Inform the group that one work is done, so the internal counter is decreased by 1.
	// It is a good idea to put defer at the top of the function because when the function
	// is long, and there may be many places with "if ... return", and we don't want to scatter
	// Done() in all those places (we may miss some places, and code looks messy).
	// Putting a defer Done() here is a neat idea.
	defer waitGroup.Done()

	mutex.Lock() // start locking this block from here to allow only ONE goroutine at a time.
	*x = *x + 1
	mutex.Unlock() // release the lock here
}

// Increment1 using a channel instead of a mutex.
// We still need to use waitGroup to tell the main thread to wait.
// Note that a channel is passed by reference, not by value in Go. So we simply pass a value for chan here.
// slices, maps, channels, pointers, functions are ALL passed by reference.
// RULE OF THUMB: anything made by make() function, new() or &, are passed to funcs by references.
func Increment1(x *int, waitGroup *sync.WaitGroup, ch chan bool) {
	defer waitGroup.Done()
	// use a channel for lock/unlocking.
	ch <- true // the channel has capacity of 1 (buffered channel)
	// any further attempts to write into the channel block until it becomes empty
	*x = *x + 1
	<-ch // read from the channel, ignore the value, so now ch is avail for writing into again.

}

func main() {
	x := 0
	// for a struct, `var waitGroup sync.WaitGroup` is equivalent to waitGroup := sync.WaitGroup{}
	// will all members of WaitGroup struct initiated to default values.
	var waitGroup sync.WaitGroup
	mutex := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1) // add one worker for each goroutine
		go Increment(&x, &waitGroup, &mutex)
	}
	waitGroup.Wait()       // block the main thread and wait for child threads to finish
	fmt.Println("x = ", x) // 1000

	// Use a channel instead of mutex to handle racing condition
	ch := make(chan bool, 1) // a buffered channel of capacity 1
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go Increment1(&x, &waitGroup, ch)
	}
	waitGroup.Wait()
	fmt.Println(x) // 2000
}
