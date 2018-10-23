package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := i; i <= 20; i++ {
		fmt.Printf("Number %d\n", i)
	}

	for j := 1; j <= 100; j++ {
		if j%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if j%3 == 0 {
			fmt.Println("Fizz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}
