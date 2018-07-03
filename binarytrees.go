package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// end of recursion
	if t == nil {
		return
	}
	// walk left recursively
	Walk(t.Left, ch)
	// send value of node to channel
	ch <- t.Value
	// walk back right recursively
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// create two channels
	ch1, ch2 := make(chan int), make(chan int)
	// create goroutine that walks first tree
	// and closes channel in the end
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	// create goroutine that walks second tree
	// and closes channel in the end
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	for {
		// read from channels with conditional ok
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		switch {
		// if values are different, stop and return false
		case v1 != v2:
			return false
		// if any of the channels is empty but not the other
		// return false because one of the trees is done first
		case ok1 != ok2:
			return false
		// if both oks are false, both trees have been traversed
		case !ok1 && !ok2:
			break
		default:
			return true
		}
	}
}

func main() {
	// we only need a 1 slot channel at a time
	ch := make(chan int, 1)
	// create goroutine that walks first tree
	// and closes channel in the end
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()
	// proof of concept, print the tree in the channel
	for i := range ch {
		fmt.Println(i)
	}
	// test if tree1 is the same as tree2
	fmt.Println("same 10 and 10 is", Same(tree.New(10), tree.New(10)))
	fmt.Println("same 1 and 2 is", Same(tree.New(1), tree.New(2)))
}
