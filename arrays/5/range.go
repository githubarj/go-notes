package main

import "fmt"

// go offers syntatic sugar to loop oveer arrays and slices

func printSliceElements (slice []string) {
	
	for index , value := range slice {
		fmt.Printf("%d : %v\n", index , value)
	}
}


func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index , m := range msg {
		for _ , b := range badWords {
			if m == b {
				return index
			}
		}
	}
	return -1
}


func main () {
	fruits := []string{"Banana", "Apple", "Orange" ,"Mango" ,"Ovacado"}

	printSliceElements(fruits)
}
