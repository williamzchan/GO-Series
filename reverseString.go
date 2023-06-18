package main

import "fmt"

func main() {
	str := "Hello, World!"
	str2 := "Hello, World! 2"
	reversed := reverseStringIterative(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversed)
	reversed2 := reverseStringRecursive(str2)
	fmt.Println("This function will use recursion to reverse the string")
	fmt.Println("Original:", str2)
	fmt.Println("Reversed:", reversed2)
}

func reverseStringIterative(input string) string {
	fmt.Println("This function will use interation to reverse the string")
	// Convert the string to a rune slice
	runes := []rune(input)

	// Get the length of the string
	length := len(runes)

	// Iterate over half the length of the string and swap characters
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	// Convert the reversed rune slice back to a string
	reversedString := string(runes)

	return reversedString
}

func reverseStringRecursive(input string) string {
	if len(input) <= 1 {
		return input
	}

	firstChar := input[0:1]
	remainingChars := input[1:]

	return reverseStringRecursive(remainingChars) + firstChar
}
