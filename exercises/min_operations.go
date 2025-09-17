package exercises

func MinimumOperations (nums [] int) int {
	operationsCounter := 0
	for _, val := range nums {
		if val % 3 != 0 {
			if (val +1) % 3 == 0 || (val -1 ) % 3 == 0 {
				operationsCounter++;
			}
		}
	} 
	return operationsCounter;
}