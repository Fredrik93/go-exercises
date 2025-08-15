package exercises

func scoreOfString(s string) int {
	arr := []byte(s)
	result := 0
	for i := 0; i < len(arr)-1; i++ {
		tempResult := int(arr[i]) - int(arr[i+1])
		result += max(tempResult, -tempResult)
	}
	return result
}
