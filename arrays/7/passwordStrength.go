package main

import (
	"fmt"
	"unicode"
)
func isValidPassword (password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	upperCase := false
	number :=false
	for _ , char := range password {
		if !upperCase && unicode.IsUpper(char) {
			upperCase = true
		} 
		if !number && unicode.IsDigit(char) {
			number = true
		}
		if upperCase && number {
			return true
		}
	}
	
	return false
}

func main () {

	fmt.Println(isValidPassword("Yjafin7ioen"))
}