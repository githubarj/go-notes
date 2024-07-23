package main

import (
	"fmt"
	"time"
)

func printPrimes(max int) {
	if max >= 2  {
		fmt.Println(2)
	}
	for n := 3 ; n <= max ; n += 2 {
		isPrime := true
		for i := 3 ; i * i <= n ; i += 2 {
			if n % i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
		
		
	}
}

// don't edit below this line

func test(max int) {
	start := time.Now() // Start timer
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Println("===============================================================")
}

func main() {
	test(3000000)
}
