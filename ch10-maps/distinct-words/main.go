package main

import "strings"

func countDistinctWords(messages []string) int {


	theMap := map[string]bool{}
	for i:=0;i<len(messages);i++ {
		loweredEntry := strings.ToLower(messages[i])
		if loweredEntry == "" {
			continue
		}
		if strings.Contains(loweredEntry, " ") {
			indivWords := strings.Split(loweredEntry, " ")
			for j:=0; j<len(indivWords);j++ {
				if _, ok := theMap[indivWords[j]]; ok {
				continue
			} else {
				theMap[indivWords[j]] = true
			}
			}

		} else {
			if _, ok := theMap[loweredEntry]; ok {
				continue
			} else {
				theMap[loweredEntry] = true
			}
		}
		
		

	}

	return len(theMap)
}

