package main

import "fmt"

func printReports(intro, body, outro string) {
	/*
	// this was wrong because it assignment asked for 3 seperate calls
	cost := 0
	message := intro + body + outro 
	introLen, bodyLen, outroLen := len(intro), len(body), len(outro)
	printCostReport(func(m string) int {
		costIntro = (len(m) - bodyLen - outroLen)*2
		costBody = (len(m) - introLen - outroLen)*3
		costIntro = (len(m) - bodyLen - outroLen)*4
		
		return cost
	}, message)
	*/

		printCostReport(func(m string) int {
		cost := len(m)*2
		return cost
	}, intro)

	printCostReport(func(m string) int {
		cost := len(m)*3
		return cost
	}, body)

	printCostReport(func(m string) int {
		cost := len(m)*4
		return cost
	}, outro)
	
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

