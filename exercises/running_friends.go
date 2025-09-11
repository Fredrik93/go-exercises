package exercises

func RecoverOrder (order [] int, friends [] int)  [] int {
	// loop over i 
	// in i loop over j (friends)
	// if j exist in i, add friends[j] to result 
	// return result 
	var result [] int;
	for _, el := range order {
		for _, friendPlace := range friends {
		 if(el == friendPlace){
			result = append(result, friendPlace)
		 }
		}	
	}

return [] int {-1}
} 