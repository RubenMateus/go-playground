package main

import (
	"fmt"
	"math"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println(greeting("Ruben"))
	fmt.Println(getSum(1, 2))
	fmt.Println(split(17))

	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// you might notice that functions with a pointer argument must take a pointer:
	var v1 Vertex
	//ScaleFunc(v1, 5)  // Compile error!
	ScaleFunc(&v1, 5) // OK

	// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	var v2 Vertex
	v1.Scale(5) // OK
	p := &v2
	p.Scale(10) // OK

	// Functions that take a value argument must take a value of that specific type:

	var v3 Vertex
	fmt.Println(AbsFunc(v3)) // OK
	//fmt.Println(AbsFunc(&v3)) // Compile error!
	//while methods with value receivers take either a value or a pointer as the receiver when they are called:

	var v4 Vertex
	fmt.Println(v.Abs()) // OK
	p2 := &v4
	fmt.Println(p2.Abs()) // OK
}
