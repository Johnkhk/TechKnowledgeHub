from visitor import Circle, Square, Triangle, AreaCalculator, PerimeterCalculator

shapes = [Circle(10), Square(5), Triangle(6, 8, 8)]
area_calculator = AreaCalculator()
perimeter_calculator = PerimeterCalculator()

for shape in shapes:
    shape.accept(area_calculator)
    shape.accept(perimeter_calculator)
