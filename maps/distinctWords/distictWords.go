package main

import "strings"

func countDistinctWords(messages []string) int {
	counter := make(map[string]int)	
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			counter[lowerWord]++

		}

	}
	return len(counter)
}
