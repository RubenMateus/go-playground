package main

import "fmt"

func main() {

	defer fmt.Println("ULTIMO")

	defer fmt.Println("PENULTIMO")

	defer fmt.Println("ANTEPENULTIMO")

	fmt.Println("PRIMEIRO")

	defer fmt.Println("SEGUNDO")

	// fmt.Println("counting....")
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
}
