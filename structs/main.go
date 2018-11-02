package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {

	person := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person1 := Person{"Samantha", "Smith", "Boston", "f", 25}

	fmt.Println(person)
	fmt.Println(person1)
}
