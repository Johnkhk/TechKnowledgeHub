from prototype import Circle, Rectangle

# Client code
circle1 = Circle(5)
circle2 = circle1.clone()
rectangle1 = Rectangle(3, 4)
rectangle2 = rectangle1.clone()

print(circle1)  # Output: Circle with radius 5
print(circle2)  # Output: Circle with radius 5
print(f"Circle area: {circle1.get_area()}")  # Output: Circle area: 78.53981633974483
print(f"Circle sides: {circle1.get_sides()}")  # Output: Circle sides: 0
print(rectangle1)  # Output: Rectangle with width 3 and height 4
print(rectangle2)  # Output: Rectangle with width 3 and height 4
print(f"Rectangle area: {rectangle1.get_area()}")  # Output: Rectangle area: 12
print(f"Rectangle sides: {rectangle1.get_sides()}")  # Output: Rectangle sides: 4
