//Package list implements a doubly linked list.

//package list


package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and populate it with numbers
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//          fmt.Println("(* sort error *)")
