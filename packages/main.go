package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/rubenmateus/go-playground/packages/stringutils"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(stringutils.Reverse("olleh"))
	fmt.Println(RepeatStr(5, "ola"))
	fmt.Println(RepeatStr1(5, "ola"))
	fmt.Println(RepeatStr2(5, "ola"))
	// lol
	//panic("aaaaaaaaaaaaaaa")

	a, b := stringutils.Swap("hello", "world")
	fmt.Println(a, b)

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

}

// simple repeat could be achieved several ways
func RepeatStr(repetitions int, value string) string {

	var str string
	for i := 0; i < repetitions; i++ {
		str += value
	}
	return str
}

func RepeatStr1(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func RepeatStr2(repetitions int, value string) string {
	if repetitions == 1 {
		return value
	}
	return RepeatStr(repetitions-1, value) + value
}
