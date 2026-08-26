package main

func getNameCounts(names []string) map[rune]map[string]int {
	// declare nested map	
	nestedMap := make(map[rune]map[string]int)

	

	// loop through names slice
	for i:=0;i<len(names);i++ {
		// for each name, create a rune slice of the letters 
		letters := []rune(names[i])
		// get first letter of ith name
		firstLetter := letters[0]

		if _, ok := nestedMap[firstLetter]; ok {

			if _,stillOk := nestedMap[firstLetter][names[i]]; stillOk {
				currVal := nestedMap[firstLetter][names[i]]
				nestedMap[firstLetter][names[i]] = currVal + 1
			} else {
				nestedMap[firstLetter][names[i]] = 1
			}


		} else {
				nestedMap[firstLetter][names[i]] = 1
		}
		
	}

	return nestedMap
}

