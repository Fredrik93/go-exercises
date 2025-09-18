package exercises

func MaximumWealth(accounts [][]int) int {
	var richestClient int
	for _,v :=range accounts {
		var tempClient int
		for _,a := range v{
			tempClient += a
		}
		if(tempClient > richestClient){ richestClient = tempClient}
	}
    return richestClient
}