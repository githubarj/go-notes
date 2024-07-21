package main 

import "fmt"


func printReports (intro, body , outro string) {
	printCostReport(func(intro string) int {
		return 2 * len(intro)
	}, intro)

	printCostReport(func(body string) int {
		return 3  * len(body)
	}, body)

	printCostReport(func(outro string) int {
		return 4 * len(outro) 
	}, outro)
}

func printCostReport (costCalc func(string) int, message string) {
	cost := costCalc(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func main() {
		printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}
