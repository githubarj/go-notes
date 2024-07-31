package main

import "fmt"

func concurrentFib(n int) []int {
	ch := make(chan int)
	arr := []int{}
	go fibonacci(n, ch)
	for item := range ch {
		arr = append(arr, item)
	}
	return arr
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
func main() {
	n := 10
	fibs := concurrentFib(n)
	fmt.Println(fibs) // Output the Fibonacci numbers
}
