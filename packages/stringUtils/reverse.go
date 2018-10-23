package stringutils

func Reverse(someString string) string {
	runeString := []rune(someString)
	var reverseString string
	for i := len(runeString) - 1; i >= 0; i-- {
		reverseString += string(runeString[i])
	}
	return reverseString
}
