package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bdasmdpas@das.co"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	names := map[string]string{"Bob": "dasdas", "sharon": "dasd"}
	fmt.Println(names)
}
