package main

import (
	"fmt"
	"github.com/githubarj/mystrings"
)

func add(a, b int) int {
	return a + b
}

func fuzzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fuzzBuzz")
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Println(i)
	}
}
func multiply(a float64, b int) float64 {
	return a * float64(b)
}
func main() {
	fmt.Println("Hello world")
	fmt.Println("I am usig neovim now")
	fmt.Println(mystrings.Reverse("Hello, world"))
	fmt.Println(add(4, 7))
	fuzzBuzz()
	fmt.Println("auto import")
	fmt.Println(multiply(45.5, 39))
}
