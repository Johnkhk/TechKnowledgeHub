package main

// Iterator interface
type Iterator interface {
	Next() *int
}

// LinkedListNode represents a node in the linked list
type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

// ForwardIterator struct
type ForwardIterator struct {
	current *LinkedListNode
}

// NewForwardIterator creates a new ForwardIterator
func NewForwardIterator(head *LinkedListNode) *ForwardIterator {
	return &ForwardIterator{current: head}
}

// Next returns the next element in forward order
func (it *ForwardIterator) Next() *int {
	if it.current == nil {
		return nil
	}
	value := it.current.Value
	it.current = it.current.Next
	return &value
}

// BackwardIterator struct
type BackwardIterator struct {
	stack []*LinkedListNode
}

// NewBackwardIterator creates a new BackwardIterator
func NewBackwardIterator(head *LinkedListNode) *BackwardIterator {
	it := &BackwardIterator{stack: []*LinkedListNode{}}
	for current := head; current != nil; current = current.Next {
		it.stack = append(it.stack, current)
	}
	return it
}

// Next returns the next element in backward order
func (it *BackwardIterator) Next() *int {
	if len(it.stack) == 0 {
		return nil
	}
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	return &node.Value
}
