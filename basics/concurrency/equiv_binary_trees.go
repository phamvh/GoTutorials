package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// https://go.dev/tour/concurrency/7
// Use channels to check if two binary trees store the same elements or not.
// This shows how easy it is to do this in Go compared to other languages.
// This also shows that two threads in GO can communicate with each other by using the same channel.
//    Normally, one writes and one receives from the same channel.

/**
This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

The function tree.New(k) constructs a randomly-structured (but always sorted)
binary tree holding the values k, 2k, 3k, ..., 10k.

We will use this func to generate trees to make use of the sorting property.
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walkRecursive(t *tree.Tree, ch chan int) {
	// in-order traversal so we get an increasingly sorted sequence
	if t.Left != nil {
		walkRecursive(t.Left, ch)
	}
	v := t.Value
	//fmt.Println(v)
	ch <- v
	if t.Right != nil {
		walkRecursive(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch) // need to close to let the receiving side know when it reaches the end.
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// create two channels of size 10 each
	ch1, ch2 := make(chan int, 10), make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for { // infinite loop
		// try to read from both channels. If trees are equivalent, two numbers should be the same.
		// these two reads block till data become avail in two channels
		// these blocking also makes the main thread wait for the two child threads to finish
		num1, ok1 := <-ch1
		num2, ok2 := <-ch2
		// we take advantage of in-order traversal here because values are sorted in increasing order for sorted trees
		if num1 != num2 || ok1 != ok2 { // if values are diff, or one channel has reached the end while the other has not
			return false
		}
		if !ok1 {
			break // both reach the end
		}

	}
	return true
}

func main() {
	t1, t2 := tree.New(1), tree.New(1) // two same trees
	fmt.Println(Same(t1, t2))
	t3 := tree.New(2)
	fmt.Println(Same(t2, t3))
}
