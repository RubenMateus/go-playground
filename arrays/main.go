package main

import "fmt"

func main() {
	// var fruitArray [2]string

	// fruitArray[0] = "apple"
	// fruitArray[1] = "orange"

	// fruitArray := [2]string{"apple", "orange"}

	fruitSlice := []string{"das", "das", ""}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

}
