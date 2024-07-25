from abc import ABC, abstractmethod


class Iterator(ABC):
    @abstractmethod
    def next(self):
        pass


class ForwardIterator(Iterator):
    def __init__(self, head):
        self.current = head

    def next(self):
        if self.current is None:
            return None
        value = self.current.value
        self.current = self.current.next
        return value


class BackwardIterator(Iterator):
    def __init__(self, head):
        self.stack = []
        current = head
        while current is not None:
            self.stack.append(current)
            current = current.next

    def next(self):
        if not self.stack:
            return None
        node = self.stack.pop()
        return node.value


class LinkedListNode:
    def __init__(self, value):
        self.value = value
        self.next = None
