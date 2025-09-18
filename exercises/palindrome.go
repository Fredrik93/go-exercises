package exercises

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {

	word := strconv.Itoa(x)
	var firstInd int = 0
	var secondInd int = len(word) - 1
	first := word[firstInd]   // byte
	second := word[secondInd] // byte

	for _, v := range word {
		fmt.Printf("first: %c, second: %c\n : v: %c", first, second, v)
		if first == second || firstInd <= secondInd {
			firstInd++
			secondInd--
		} else {
			return false
		}

	}
	return false
}
