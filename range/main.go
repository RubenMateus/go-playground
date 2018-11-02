package main

import "fmt"

func main() {
	ids := []int{33, 76, 56, 65, 4651, 123, 132, 8478, 4, 6, 6}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	names := map[string]string{"Bob": "dasdas", "sharon": "dasd"}
	for k, v := range names {
		fmt.Printf("%s: %s", k, v)
	}
}
