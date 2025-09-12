package exercises

func RecoverOrder (order [] int, friends [] int)  [] int {
	var result [] int;
	for _, el := range order {
		for _, friendPlace := range friends {
		 if(el == friendPlace){
			result = append(result, friendPlace)
		 }
		}	
	}

return result
} 