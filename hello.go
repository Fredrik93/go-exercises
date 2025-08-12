package main

import (
	"fmt"
	"unicode"
)

func ReverseWordsPreserveStructure(input string) string {
	// convert string to char arr
	runes := []rune(input)
	stack := []rune{}
	var result []rune = []rune{}
	// iterate through the runes and reverse each word
	for i := 0; i < len(runes); i++ {
		var symbol rune = runes[i];
		fmt.Printf("symbol (character): %c\n", symbol)
		if unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]) {
			stack = append(stack, runes[i])

		} else if unicode.IsSpace(runes[i]) {

			for len(stack) > 0 {

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = append(result, top)
			}
			result = append(result, runes[i]) // add the space or punctuation

		}

	}
	for len(stack) > 0 {

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = append(result, top)
			}

	return string(result)
}

func main() {
}
