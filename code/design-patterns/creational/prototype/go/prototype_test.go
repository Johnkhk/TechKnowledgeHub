package prototype

import "fmt"

// Client code
func main() {
	circle1 := &Circle{Radius: 5}
	circle2 := circle1.Clone()
	rectangle1 := &Rectangle{Width: 3, Height: 4}
	rectangle2 := rectangle1.Clone()

	fmt.Println(circle1)                                       // Output: Circle with radius 5
	fmt.Println(circle2)                                       // Output: Circle with radius 5
	fmt.Printf("Circle area: %f\n", circle1.GetArea())         // Output: Circle area: 78.539816
	fmt.Printf("Circle sides: %d\n", circle1.GetSides())       // Output: Circle sides: 0
	fmt.Println(rectangle1)                                    // Output: Rectangle with width 3 and height 4
	fmt.Println(rectangle2)                                    // Output: Rectangle with width 3 and height 4
	fmt.Printf("Rectangle area: %f\n", rectangle1.GetArea())   // Output: Rectangle area: 12.000000
	fmt.Printf("Rectangle sides: %d\n", rectangle1.GetSides()) // Output: Rectangle sides: 4
}
