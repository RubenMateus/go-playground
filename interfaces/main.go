package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle Area: %f\n", GetArea(circle))
	fmt.Printf("Rectangle Area: %f\n", GetArea(rectangle))
}
