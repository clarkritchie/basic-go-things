package main

import (
	"fmt"
	"math"
)

// Define a basic interface named Shape
type Shape interface {
	area() float64
}

// Define a struct named Circle
type Circle struct {
	radius float64
}

// Define a method for Circle to calculate its area
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Define a struct named Rectangle
type Rectangle struct {
	width, height float64
}

// Define a method for Rectangle to calculate its area
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 3, height: 4}

	// Call a function that accepts Shape interface
	printArea(circle)
	printArea(rectangle)
}

// Function to accept any Shape and print its area
func printArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.area())
}
