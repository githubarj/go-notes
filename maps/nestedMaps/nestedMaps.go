package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	if len(names) > 0 {
		for _, name := range names {
			chars := []rune(name)
			firstChar := chars[0]

			if _ , ok :=  result[firstChar]; !ok {
				result[firstChar] = make(map[string]int)
			}

			result[firstChar][name]++
		}
	}
	return result
}

func main() {
	names := []string{"Alice", "Bob", "Bob",  "Charlie", "David", "Alvin", "Ben", "Carol"}
	result := getNameCounts(names)
	fmt.Println(result)
}