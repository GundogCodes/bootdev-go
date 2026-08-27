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
		// if firstLetter key exists
		if _, ok := nestedMap[firstLetter]; ok {
			// if name key already exists
			if _,stillOk := nestedMap[firstLetter][names[i]]; stillOk {
				//get value
				currVal := nestedMap[firstLetter][names[i]]
				// increment value
				nestedMap[firstLetter][names[i]] = currVal + 1
			} else {
				// if name key does not exist add key and set value to one
				nestedMap[firstLetter][names[i]] = 1
			}

		// if firstLetter key does not exist add
		} else {
				// make an inner map
				innerMap := make(map[string]int)
				// set name and value
				innerMap = [names[i]] = 1
				// set inner map as value to nestedMap
				nestedMap[firstLetter] = innerMap
		}
		
	}

	return nestedMap
}
