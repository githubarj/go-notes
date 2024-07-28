package main

import (
	"fmt"
)

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _ , ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}


func main() {
	
	messagedUsers := []string{"benji", "alice", "benji", "bob", "benji", "alice", "charlie"}

	
	validUsers := map[string]int{
		"benji":  0,
		"alice":  0,
		"bob":    0,
		"dave":   0,
	}

	getCounts(messagedUsers, validUsers)

	for user, count := range validUsers {
		fmt.Printf("User: %s, Message Count: %d\n", user, count)
	}
}