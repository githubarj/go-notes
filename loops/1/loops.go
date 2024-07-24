package main

import "fmt"

func bulkSend (numMessages int) float64 {
	totalCost := 0.0
	for i:= 0; i < numMessages ; i ++ {
		messageCost := 1.0 + (float64(i) * 0.01)
		totalCost += messageCost
	}
	return totalCost
}

func maxMessages (thresh int) int {
	
	totalCost := 0

	for i := 0 ;  ; i ++ {
		totalCost += 100 + i
		if totalCost > thresh {
			return i
		}
	}
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxCostInPennies)  {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies >= float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func fizzBuzz () {
	for i:= 1 ; i <= 100 ; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main () {
	fmt.Println(bulkSend(20))
	fmt.Println(maxMessages(400))
	fmt.Println(getMaxMessagesToSend(1.1 , 5))
	fizzBuzz()
}