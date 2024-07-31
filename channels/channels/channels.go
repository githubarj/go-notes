package main

import (
	"fmt"
	"time"
)

type email struct {
  body string
  date time.Time
}

func checkEmailAge (emails [3] email) [3] bool{
  isOldChan := make(chan bool , len(emails))

  go sendIsOld(isOldChan, emails)

  isOld := [3]bool{}
  isOld[0] = <- isOldChan
  isOld[1] = <- isOldChan
  isOld[2] = <- isOldChan

  close(isOldChan)
  return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func main() {
	emails := [3]email{
		{body: "Email 1", date: time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC)},
		{body: "Email 2", date: time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC)},
		{body: "Email 3", date: time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)},
	}

	isOld := checkEmailAge(emails)
	fmt.Println(isOld)
}
