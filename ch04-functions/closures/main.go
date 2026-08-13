package main


// A closure is a function that references variables from outside its own function body.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

