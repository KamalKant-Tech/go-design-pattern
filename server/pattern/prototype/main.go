package main

import "fmt"

type ShapeType int // Custom type

const (
	SquareType ShapeType = 1
	CircleType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType // get the shape id
	Clone() Shape     // clone the shape object or get the copy of the objects
	PrintProp()       // Print the type of the properties
}

type Circle struct {
	Id            ShapeType
	Radius        int
	Diameter      int
	Circumference int
}

type Square struct {
	Id      ShapeType
	Length  int
	Breadth int
}

// Lets implement the shape interface for circle
func (c *Circle) GetId() ShapeType {
	return c.Id
}

// Create the circle objects
func NewCircle(radius, diameter, circumference int) Shape {
	fmt.Println("Inside clone function")
	return &Circle{CircleType, radius, diameter, circumference}
}

// Implement clone method in circle struct
func (c *Circle) Clone() Shape {
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}

// Implement PrintTop interface function in Circle
func (c *Circle) PrintProp() {
	fmt.Println("Circle Properties Radius, Diameter, Circumference:", c.Radius, c.Diameter, c.Circumference)
}

func NewSquare(Length, Breadth int) Shape {
	return &Square{SquareType, Length, Breadth}
}

func (s *Square) GetId() ShapeType {
	return s.Id
}

func (s *Square) Clone() Shape {
	return NewSquare(s.Length, s.Breadth)
}

func (s *Square) PrintProp() {
	fmt.Println("Square Properties Length, Breadth: ", s.Length, s.Breadth)
}

var RegisterShapes = make(map[int]*Shape) // Prototype registery

func LoadRegistery() {
	circle := NewCircle(40, 30, 20)
	RegisterShapes[int(circle.GetId())] = &circle // adding circle object to registery

	square := NewSquare(40, 30)
	RegisterShapes[int(square.GetId())] = &square // adding square object to registery

}

func main() {
	LoadRegistery()

	circle := *RegisterShapes[int(CircleType)]
	if c, ok := circle.(*Circle); ok {
		fmt.Println("Older Properties")
		c.PrintProp()

		cClone := c.Clone().(*Circle)
		cClone.Radius = 30
		fmt.Println("New Properties with clone objects")
		fmt.Printf("Type of %v is %T", c, c)
		cClone.PrintProp()
	}

	square := *RegisterShapes[int(SquareType)]
	if c, ok := square.(*Square); ok {
		fmt.Println("Older Properties")
		c.PrintProp()

		cClone := c.Clone().(*Square)
		cClone.Breadth = 70
		fmt.Println("New Properties with clone objects")
		//fmt.Printf("Type of %v is %T", c, c)
		cClone.PrintProp()
	}
}
