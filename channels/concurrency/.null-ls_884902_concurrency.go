package main

import (
  "fmt"
  "time"
)

func sendMail (message string) {
  go func (){
    time.Sleep(time.Millisecond * 250)
    fmt.Println("Email received: '%s' \n", message)
  }() 
    fmt.Printf("Email sent: '%s'\n", message)
  
}

func test (message string) {
  sendMail(message)
  time.Sleep(time.Millisecond * 500)
  fmt.Println("=================================")
}

func main () {
  test("Hello there Kaladin")
  test("Hey there Sharon")
  test("Hi there Sharon")
}
