package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i:=0; i<len(msg); i++ {
		potWord :=  msg[i]
		for j:=0; j<len(badWords); j++{
			if potWord == badWords[j] {
				return i
			}
		}
	}
	return -1
}

