package prototype

import (
	"fmt"
	"math"
)

// Prototype interface
type Shape interface {
	Clone() Shape
	GetArea() float64
	GetSides() int
}

// Concrete prototype
type Circle struct {
	Radius int
}

func (c *Circle) Clone() Shape {
	// Create a new instance of Circle and copy all fields
	return &Circle{Radius: c.Radius}
}

func (c *Circle) GetArea() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c *Circle) GetSides() int {
	return 0
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle with radius %d", c.Radius)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) Clone() Shape {
	// Create a new instance of Rectangle and copy all fields
	return &Rectangle{Width: r.Width, Height: r.Height}
}

func (r *Rectangle) GetArea() float64 {
	return float64(r.Width * r.Height)
}

func (r *Rectangle) GetSides() int {
	return 4
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle with width %d and height %d", r.Width, r.Height)
}
