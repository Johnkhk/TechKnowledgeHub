from abc import ABC, abstractmethod
import math


# Prototype interface
class Shape(ABC):
    @abstractmethod
    def clone(self):
        pass

    @abstractmethod
    def get_area(self):
        pass

    @abstractmethod
    def get_sides(self):
        pass


# Concrete prototype
class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def clone(self):
        # Create a new instance of Circle and copy all fields
        new_circle = Circle(self.radius)
        return new_circle

    def get_area(self):
        return math.pi * self.radius**2

    def get_sides(self):
        return 0

    def __str__(self):
        return f"Circle with radius {self.radius}"


class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def clone(self):
        # Create a new instance of Rectangle and copy all fields
        new_rectangle = Rectangle(self.width, self.height)
        return new_rectangle

    def get_area(self):
        return self.width * self.height

    def get_sides(self):
        return 4

    def __str__(self):
        return f"Rectangle with width {self.width} and height {self.height}"
