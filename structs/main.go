package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
	isMarried bool

	// or
	// firstName1, lastName1, city1, gender1 string
	// age1                                  int
	// etc...
}

func (person Person) Greet() string {
	return "Hello, my name is " + person.firstName + " " +
		person.lastName + " with " + strconv.Itoa(person.age) + " and Married: " + strconv.FormatBool(person.isMarried)
}

func (person *Person) HasBirthday() {
	person.age++
}

func (person *Person) GetMarried(spouseLastName string) {
	if person.gender == "f" {
		person.lastName = spouseLastName
	}

	person.isMarried = true
}

func main() {

	person := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person1 := Person{"Samantha", "Smith", "Boston", "f", 25, false}

	fmt.Println(person)
	fmt.Println(person1)

	person.HasBirthday()
	person.GetMarried(person1.lastName)
	fmt.Println(person.Greet())

}
