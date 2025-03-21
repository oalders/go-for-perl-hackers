// package main implements a linked list courtesy of ChatGPT.
package main

import "fmt"

// LinkedList represents the entire linked list with a reference to the head node.
type LinkedList struct {
	Head *Node
}

// Node represents a single node in the linked list.
type Node struct {
	Next    *Node // Pointer to the next node in the list
	Payload int   // Can be any type depending on your requirements
}

// Add adds a new node to the end of the list.
func (ll *LinkedList) Add(data int) {
	newNode := &Node{Payload: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Print prints the entire linked list.
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Payload)
		current = current.Next
	}
	fmt.Println("nil")
}

// Main function to test the linked list.
func main() {
	ll := &LinkedList{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Print() // Output: 10 -> 20 -> 30 -> nil
}
