from iterator import ForwardIterator, BackwardIterator, LinkedListNode

head = LinkedListNode(1)
head.next = LinkedListNode(2)
head.next.next = LinkedListNode(3)
head.next.next.next = LinkedListNode(4)
head.next.next.next.next = LinkedListNode(5)

# Forward traversal
forward_iterator = ForwardIterator(head)
print("Forward traversal: ", end="")
value = forward_iterator.next()
while value is not None:
    print(value, end=" ")
    value = forward_iterator.next()
print()

# Backward traversal
backward_iterator = BackwardIterator(head)
print("Backward traversal: ", end="")
value = backward_iterator.next()
while value is not None:
    print(value, end=" ")
    value = backward_iterator.next()
print()
