package main

import "fmt"

func main() {
	//var name = "Ruben"
	var age int32 = 28
	const isCool = true
	name := "Ruben"
	size := 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)

}
