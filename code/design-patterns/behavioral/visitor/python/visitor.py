from abc import ABC, abstractmethod


# Visitor interface
class Visitor(ABC):
    @abstractmethod
    def visit_circle(self, circle):
        pass

    @abstractmethod
    def visit_square(self, square):
        pass

    @abstractmethod
    def visit_triangle(self, triangle):
        pass


# Concrete Visitor for calculating area
class AreaCalculator(Visitor):
    def visit_circle(self, circle):
        area = 3.14 * (circle.radius**2)
        print(f"Circle area: {area}")

    def visit_square(self, square):
        area = square.side * square.side
        print(f"Square area: {area}")

    def visit_triangle(self, triangle):
        area = 0.5 * triangle.base * triangle.height
        print(f"Triangle area: {area}")


# Concrete Visitor for calculating perimeter
class PerimeterCalculator(Visitor):
    def visit_circle(self, circle):
        perimeter = 2 * 3.14 * circle.radius
        print(f"Circle perimeter: {perimeter}")

    def visit_square(self, square):
        perimeter = 4 * square.side
        print(f"Square perimeter: {perimeter}")

    def visit_triangle(self, triangle):
        perimeter = triangle.side_a + triangle.side_b + triangle.side_c
        print(f"Triangle perimeter: {perimeter}")


# Shape interface
class Shape(ABC):
    @abstractmethod
    def accept(self, visitor):
        pass


# Concrete Shape: Circle
class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def accept(self, visitor):
        visitor.visit_circle(self)


# Concrete Shape: Square
class Square(Shape):
    def __init__(self, side):
        self.side = side

    def accept(self, visitor):
        visitor.visit_square(self)


# Concrete Shape: Triangle
class Triangle(Shape):
    def __init__(self, side_a, side_b, side_c):
        self.side_a = side_a
        self.side_b = side_b
        self.side_c = side_c
        self.base = side_a
        self.height = (
            side_b**2 - (0.5 * side_a) ** 2
        ) ** 0.5  # assuming it's an isosceles triangle

    def accept(self, visitor):
        visitor.visit_triangle(self)
