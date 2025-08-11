package main

import "fmt"

func ReverseWordsPreserveStructure(input string) string {
	// convert string to char arr
	runes := []rune(input)
	// iterate through the runes and reverse each word
	for i := 0; i < len(runes); i++ {
		fmt.Printf("rune: %c\n", runes[i])
	}
	// reverse each word while preserving spaces and punctuation
	// return the modified string

	return input
}

func main() {
	fmt.Println("Hello, World!")
}
