package exercises

func FindClosest(x int, y int, z int) int {

	var firstPersonDist int = absoluteValue(x, z)
	var secondPersonDist int = absoluteValue(y, z)

	if firstPersonDist < secondPersonDist {
		return 1
	} else if firstPersonDist > secondPersonDist {
		return 2
	} else if firstPersonDist == secondPersonDist {
		return 0
	}
	return -1
}
func absoluteValue(firstValue int, secondValue int) int {
	value := firstValue - secondValue
	if value < 0 {
		return value * -1
	}
	return value
}
