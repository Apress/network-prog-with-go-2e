package main

import (
	"fmt"
	"sync"
	"time"
)

// Our LinkedList code
type LL struct {
	N    *LL
	data string
}

// Retreive the next node in the linkedlist
func (l LL) Next() *LL { return l.N }

// Here we use a union of types “|”
// Meaning our arguments must be of these types
// Either channel  or the above linkedlist type
type MustBe interface {
	chan string | LL
}

// We use the union technique once more on the return types
// Notice these can differ than the above "MustBe" types
type Result interface {
	string | LL
}

// This is the function we wish to make generic
// We are iterating over a instance of MustBe (either channel string or LL)
// Notice the return type must be a Result type (either string or LL)
func Iterate[M MustBe, R Result](o M, iter func(M) R) (r R) {
	return iter(o)
}

func main() {
	// Create channel for strings
	c := make(chan string, 5)
	c <- "ok"
	c <- "ok2"

	//This function is what we will pass into the above Iterate
	//Notice Iterate's first parameter is the same as the following
	//lambdas parameter
	citer := func(c chan string) string {
		select {
		case msg1 := <-c:
			return msg1
		case <-time.After(1 * time.Second):
			return "nothing"
		}
	}

	// Here we "Iterate" through the channel
	var wg sync.WaitGroup
	wg.Add(2)
	go func(f func(chan string) string) {
		for {
			fmt.Println(Iterate(c, f))
			wg.Done()
		}
	}(citer)
	wg.Wait() // wait for iteration to finish

	// The remaining example shows passing a custom Linked List
	// iteration function

	// First we build a simple list
	n1 := LL{data: "n1"}
	n2 := LL{data: "n2"}
	n3 := LL{data: "n3"}
	n1.N = &n2
	n2.N = &n3

	// Like the above citer, the parameter type will match
	// the first parameter of Iterate above
	liter := func(l LL) LL {
		var zero LL

		if l.N != zero.N {
			return *l.N
		} else {
			return zero
		}
	}

	// We walk through the linked list
	n := n1
	for n.N != nil {
		fmt.Printf("node:%s\n",n.data)
		n = Iterate(n, liter)
	}
}
