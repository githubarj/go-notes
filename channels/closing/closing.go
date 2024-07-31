package main

import "fmt"

func countReports(numSentCh chan int) int {
  count := 0
	for {
    numReports, ok := <- numSentCh
    if !ok {
      break
    }
    count = numReports
  }
  return count
}


func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}


func main() {
    numBatches := 5
    numSentCh := make(chan int)

    // Start the goroutine to send reports
    go sendReports(numBatches, numSentCh)

    // Count reports and print the last received value
    result := countReports(numSentCh)
    fmt.Printf("Last received count: %d\n", result)
}
