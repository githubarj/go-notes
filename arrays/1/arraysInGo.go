package main

import (
	"errors"
	"fmt"
)

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{}

	costs[0] = len(primary)
	costs[1] = costs[0] + len(secondary)
	costs[2] = costs[1] + len(tertiary)

	return messages, costs
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case "pro":
		return messages[:] , nil
	case "free":
		return messages[:2] , nil
	default:		
		return nil , errors.New("unsupported plan")
	}
}

func getMessagesCost (messages []string) [] float64 {
	messagesCost := make([] float64 , len(messages))
	for i := 0 ; i < len(messages)  ;i++ {
		messagesCost[i] = float64(len(messages[i])) * 0.01
	}
	return messagesCost
}

func main() {
	primary := "hello"
	secondary := "world"
	tertiary := "!"
	
	messages, costs := getMessageWithRetries(primary, secondary, tertiary)

	fmt.Println("Messages:", messages)
	fmt.Println("Costs:", costs)

		newMessages := [3]string{"Message 1", "Message 2", "Message 3"}
	
	result, err := getMessageWithRetriesForPlan("pro", newMessages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Messages:", result)
	}
	
	result, err = getMessageWithRetriesForPlan("free", newMessages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Messages:", result)
	}
	
	result, err = getMessageWithRetriesForPlan("unknown", newMessages)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Messages:", result)
	}
}
