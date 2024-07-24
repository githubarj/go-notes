package main 

import "fmt"

func concat (strs ...string) string {
	final := ""

	for i := 0 ; i < len(strs) ; i ++ {
		final += strs[i] + " "
	}
	return final
}

// the spread operator allows us to pass a slice into a variadic function
func printThreeStrings (strs ...string) {
	for i := 0; i < len(strs) ; i++ {
		fmt.Println(strs[i])
	}
}

func sum (nums ...int) int {
	total := 0
	for i:= 0; i < len(nums) ; i++ {
		total += nums[i]
	}
	return total
}

func main () {
	final := concat("Hello", "there", "friend")
	fmt.Println(final)

	// passing a slice to a variadic function looks a bit different 
	names := []string{"John", "Mary", "Julie"}

	printThreeStrings(names...)
}