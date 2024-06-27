from abc import ABC, abstractmethod
from typing import List


class Component(ABC):
    @abstractmethod
    def calculate_size(self) -> int:
        pass


class File(Component):
    def __init__(self, name: str, size: int):
        self.name = name
        self.size = size

    def calculate_size(self) -> int:
        return self.size


class Directory(Component):
    def __init__(self, name: str):
        self.name = name
        self.children: List[Component] = []

    def add(self, component: Component):
        self.children.append(component)

    def calculate_size(self) -> int:
        total_size = 0
        for child in self.children:
            total_size += child.calculate_size()
        return total_size
