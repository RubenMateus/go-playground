package main

import "fmt"

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

func main() {
	fmt.Println(greeting("Ruben"))
	fmt.Println(getSum(1, 2))
	fmt.Println(split(17))
}
