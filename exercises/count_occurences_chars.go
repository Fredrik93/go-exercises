package exercises

import (
	"strings"
)

func NumJewelsInStones(jewels string, stones string) int {
	var result int
    for i, s := range stones {
		if(strings.ContainsRune(jewels, s)){
			i++
			result++
		}
	}
	
	return result

}