package exercises

func Shuffle(nums []int, n int) []int {
	var result []int
	//:n half of the list (first half)
	var first []int = nums[:n]
	//n: half of the list (second half)
	var second []int = nums[n:]
		for j := 0; j < n; j++ {

			result = append(result, first[j])
			result = append(result, second[j])
		}
	return result
}
