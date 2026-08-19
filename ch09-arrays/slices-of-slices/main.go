package main

func createMatrix(rows, cols int) [][]int {
	mat := [][]int{}
	for i:=0; i<rows; i++ {
		miniSlice := []int{}
		for j:=0; j<cols; j++ {
			
			miniSlice = append(miniSlice, i*j)
		}
		mat = append(mat, miniSlice)
	}
	return mat
}


// You can't index into a slice that has no elements yet.

// don't do this!
// someSlice = append(otherSlice, element)
// The append() function only creates a new array when there isn't any capacity left. 