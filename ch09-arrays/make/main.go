package main

// Most of the time we don't need to think about
// the underlying array of a slice. We can create 
// a new slice using the make function:
// func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)


func getMessageCosts(messages []string) []float64 {
	lenOfMessages := len(messages)
	costSlice := make([]float64, lenOfMessages)
	for i:=0; i<lenOfMessages; i++ {
		costSlice[i] = float64(len(messages[i])) * 0.01
	}

	return costSlice
	
}
