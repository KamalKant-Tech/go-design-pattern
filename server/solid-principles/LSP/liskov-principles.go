package main

// Liskov Substitution Principle says that an object of super class can be replaceble by an object of sub class without breaking the correctness of the program.
// In this below example if we observe and can see that Square in super class here and circle is subclass. So if I do not have the area in subclass or implement the super class area function call then its violating the LSP. This is happening cause area of circle returns the wrong value.
// To overcome this problem I have overide the Area() Parent into its child class. Now, its follow the LSP and return the correct values for respective types.
import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Width float64
}

type Circle struct {
	Square
	Radius float64
}

func (s *Square) Area() float64 {
	return s.Width * s.Width
}

// To violate the LSP principle here will return the area of square.
// This will return wrong value area of circle which against the LSP.
func (c *Circle) Area() float64 {
	//return c.Square.Area()
	return math.Pi * c.Radius * c.Radius
}

// Good example: A function that works with any shape that implements the Shape interface

func calculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	s := &Square{10}
	fmt.Println(calculateArea(s))

	c := &Circle{Radius: 5}
	fmt.Println(calculateArea(c))
}
