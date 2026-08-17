package main
// connection size is probably groupSize * groupSize - 1
func countConnections(groupSize int) int {
	return groupSize * (groupSize - 1) / 2
}

// gs = 2, c = 1
// gs = 3, c 