package main

import (
	"fmt"
	"math"
)

// Visitor interface
type Visitor interface {
	VisitCircle(c *Circle)
	VisitSquare(s *Square)
	VisitTriangle(t *Triangle)
}

// Concrete Visitor for calculating area
type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("Circle area: %.2f\n", area)
}

func (a *AreaCalculator) VisitSquare(s *Square) {
	area := s.Side * s.Side
	fmt.Printf("Square area: %.2f\n", area)
}

func (a *AreaCalculator) VisitTriangle(t *Triangle) {
	area := 0.5 * t.Base * t.Height
	fmt.Printf("Triangle area: %.2f\n", area)
}

// Concrete Visitor for calculating perimeter
type PerimeterCalculator struct{}

func (p *PerimeterCalculator) VisitCircle(c *Circle) {
	perimeter := 2 * math.Pi * c.Radius
	fmt.Printf("Circle perimeter: %.2f\n", perimeter)
}

func (p *PerimeterCalculator) VisitSquare(s *Square) {
	perimeter := 4 * s.Side
	fmt.Printf("Square perimeter: %.2f\n", perimeter)
}

func (p *PerimeterCalculator) VisitTriangle(t *Triangle) {
	perimeter := t.SideA + t.SideB + t.SideC
	fmt.Printf("Triangle perimeter: %.2f\n", perimeter)
}

// Shape interface
type Shape interface {
	Accept(visitor Visitor)
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Square struct
type Square struct {
	Side float64
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Triangle struct
type Triangle struct {
	SideA  float64
	SideB  float64
	SideC  float64
	Base   float64
	Height float64
}

func NewTriangle(sideA, sideB, sideC float64) *Triangle {
	// assuming it's an isosceles triangle for simplicity
	height := math.Sqrt(sideB*sideB - (0.5*sideA)*(0.5*sideA))
	return &Triangle{
		SideA:  sideA,
		SideB:  sideB,
		SideC:  sideC,
		Base:   sideA,
		Height: height,
	}
}

func (t *Triangle) Accept(visitor Visitor) {
	visitor.VisitTriangle(t)
}
