package main

func main() {
	shapes := []Shape{
		&Circle{Radius: 10},
		&Square{Side: 5},
		NewTriangle(6, 8, 8),
	}

	areaCalculator := &AreaCalculator{}
	perimeterCalculator := &PerimeterCalculator{}

	for _, shape := range shapes {
		shape.Accept(areaCalculator)
		shape.Accept(perimeterCalculator)
	}
}
