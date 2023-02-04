package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Mutex provides lock and unlock functions to clock a block of code.
After calling Lock(), only one goroutine can access a block of code.

We can define a block of code to be executed in mutual exclusion by
surrounding it with a call to Lock and Unlock.

We can also use defer to ensure the mutex will be unlocked
*/

// SafeCounter is data struct that allows ony one goroutine to update/get data at a time.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc method increments the value of a key in the map by one.
func (c *SafeCounter) Inc(key string) {
	// Note: I tried removing the lock here, and got err: fatal error: concurrent map writes
	c.mu.Lock() // start locking this block
	// now, only one goroutine can access this part of code from this point
	c.v[key]++    // I guess we can do this because c is a pointer, but weird syntax.
	c.mu.Unlock() // unlock - the end of the sync block
}

// Value func returns the value of a give key.
func (c *SafeCounter) Value(key string) int {
	// I don't understand the point of locking here cuz it's reading only
	// I tried removing the lock here and make 5 goroutines to call this function and it was fine.
	c.mu.Lock()
	// the retrieval from a map can throw an error if key does not exist (b/c we do not use val, ok := c.v[key] here
	// but we want to make sure that we unlock this, so we use defer
	defer c.mu.Unlock()
	return c.v[key] // the Unlock will execute after this return
}

func main() {
	c := SafeCounter{
		v: make(map[string]int),
	}
	for i := 0; i < 5; i++ {
		go c.Inc("mykey")
	}
	//time.Sleep(time.Second) // sleep to make sure we have time to run 5 updates in the loop above.
	for i := 0; i < 5; i++ {
		go fmt.Println(c.Value("mykey"))
	}
	time.Sleep(time.Second)
	//fmt.Println(c.Value("mykey"))
}
