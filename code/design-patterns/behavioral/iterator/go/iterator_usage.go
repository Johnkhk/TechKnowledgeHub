package main

import "fmt"

func main() {
	// Create a linked list
	head := &LinkedListNode{Value: 1}
	head.Next = &LinkedListNode{Value: 2}
	head.Next.Next = &LinkedListNode{Value: 3}
	head.Next.Next.Next = &LinkedListNode{Value: 4}
	head.Next.Next.Next.Next = &LinkedListNode{Value: 5}

	// Forward traversal
	forwardIt := NewForwardIterator(head)
	fmt.Print("Forward traversal: ")
	for value := forwardIt.Next(); value != nil; value = forwardIt.Next() {
		fmt.Print(*value, " ")
	}
	fmt.Println()

	// Backward traversal
	backwardIt := NewBackwardIterator(head)
	fmt.Print("Backward traversal: ")
	for value := backwardIt.Next(); value != nil; value = backwardIt.Next() {
		fmt.Print(*value, " ")
	}
	fmt.Println()
}
