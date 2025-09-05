package exercises

import "sort"

func TransformArray(nums []int) []int {
	result := []int{}

	//loop over nums
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}

	}
	sort.Ints(result)
	
	return result
}
