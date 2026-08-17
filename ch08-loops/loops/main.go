package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.0 //1
	inc := 0.00 
	for i := 0; i < numMessages; i++ {
		totalCost = totalCost + 1.00 + inc
		inc = inc + 0.01
	}
	return totalCost
}
