package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	cost1 := len(primary)
	cost2 := len(primary) + len(secondary)
	cost3 := len(primary) + len(secondary) + len(tertiary)
	m := [3]string{primary, secondary, tertiary}
	c := [3]int{cost1, cost2, cost3}
	return m, c 
}

