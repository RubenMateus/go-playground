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

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	//A type implements an interface by implementing its methods. There is no explicit declaration of intent, no //"implements" keyword.
	//Implicit interfaces decouple the definition of an interface from its implementation, which could then appear //in any package without prearrangement.

	fmt.Printf("Circle Area: %f\n", GetArea(circle))
	fmt.Printf("Rectangle Area: %f\n", GetArea(rectangle))

	var i I = T{"hello"}
	i.M()

	//An empty interface may hold values of any type. (Every type implements at least zero methods.)
	var inter interface{}
	describe(i)

	inter = 42
	describe(inter)

	inter = "hello"
	describe(inter)

	var cena interface{} = "hello"

	s := cena.(string)
	fmt.Println(s)

	s, ok := cena.(string)
	fmt.Println(s, ok)

	f, ok := cena.(float64)
	fmt.Println(f, ok)

	//f = cena.(float64) // panic
	//fmt.Println(f)

	do(21)
	do("hello")
	do(true)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
