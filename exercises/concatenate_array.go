package exercises

func GetConcatenation(nums []int) []int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	nums = append(nums, numsCopy...)
	return nums
}
