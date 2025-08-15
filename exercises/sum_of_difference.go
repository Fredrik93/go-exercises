package exercises


func DifferenceOfSum(n int, m int) int {
	divisible := 0
	notDivisible := 0
	for i := 0; i < n+1; i++ {
		if i%m == 0 {
			divisible += i
		} else {
			notDivisible += i
		}
	}
	return notDivisible - divisible
}