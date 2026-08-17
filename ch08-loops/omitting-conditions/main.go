package main

func maxMessages(thresh int) int {
	numOfMessages := 0
	cost := 100
	for i := 0; ; i++ {
		thresh = thresh - cost - i
		if thresh < 0 {
			return numOfMessages
		}
		numOfMessages += 1
	}

}

// bruh this too me way longer than it needed to