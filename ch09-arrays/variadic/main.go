package main
// Variadic just means "accept any number of arguments".

// The spread operator is the reverse when you already 
// have a slice and want to pass it into a variadic function, 
// you use ... to unpack it.
func sum(nums ...int) int {
	total := 0
	len := len(nums)
	for i:=0; i<len; i++ {
		total  = total + nums[i]
	}
	return total
}

