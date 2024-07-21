package main

import "fmt"
// if all the arguments of a function are of one type go allows us to put the type only once 
// after the last argument
func addTwo (a , b int) int {
	return a + b
}
func concat (str1 string, str2 string) string  {
	return str1 + str2
}

func getBillForMonth(cost , messages int) int {
	return cost * messages
}

func monthlyBillIncrease (costPerSpend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSpend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSpend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func returnTwo (x, y int) (int, int) {
	return x, y
}

func getNames() (string, string) {
	return "John" , "Doe"
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking,yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking , yearsUntilCarRental
}

func getAmount(num int) string {
	if num >= 90 {
		return "A" 
	}
	if num >= 70 {
		return "B" 
	}
	if num >= 50 {
		return "C"
	}
	return "D"
}

func main(){
	fmt.Println(getAmount(78))
	fmt.Println(getAmount(34))
	fmt.Println(yearsUntilEvents(17))
	firstName,_ := getNames()
	fmt.Println("Hello and how are you", firstName)
	fmt.Println(returnTwo(4,5))
	fmt.Println(monthlyBillIncrease(2, 45, 59))
	fmt.Println(addTwo(1,2))
	fmt.Println(concat("Elaine", "Mwihoko"))
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string
	congrats := "happy birthday"
	fmt.Printf(
		"%v %f %v %q %s \n", 
		smsSendingLimit,
		costPerSms,
		hasPermission,
		username,
		congrats,
	)

	penniesPerText := 2.0
	fmt.Printf("The type of this variable is %T\n ",penniesPerText)
	
	averageOpenRate , displayMessage := 0.23 , "is the average open rate of your messages"
	fmt.Println( averageOpenRate, displayMessage)

	accountAge := 2.6 
	accountAgeInt := int(accountAge)

	fmt.Println(accountAgeInt)
	const premiumPlan string  = "Premium Plan"
        const basicPlan = "Basic plan"
	fmt.Println("Plan: ", premiumPlan)
	fmt.Println("Plan: ", basicPlan)

	const secondsInaMinute = 60
	const minutesInanHour = 60
	const secondsInHour = secondsInaMinute * minutesInanHour
	fmt.Println(secondsInHour)

	const name = "Richard Jeremy"
	const openRate  = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n ",name, openRate ) 
	fmt.Println(msg)

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length", messageLen)
	
	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	}else {
		fmt.Println("Message not sent")
	}
	
}

// types in go 
// bool , string 
// number types in go - int , int8 , int16, int32 , int64 sized numbers
// unsigned numbers in go - uint,  uint8 , uint16 , uint32 uint64,  uintptr
// rune - represents a unicode code point - used to rep one character in a string , alias for the int32
// byte - alias for an unsigned 8 bit integer , uint8 - can integer values between 0 and 8
// float32 and float64 - this store floating pont numbers, float32 has up to 7 decimal point of precision 
// float64 has up to 15 or 16 , and is the default inferred type 
// complex64 and complex128 - store complex numbers 
// nibble is 4 bits while byte is 8 bits 
// go supports the const keyword which is used to create immutable values which cannot be changed after declaration
// fmt.Printf - %v (default) , %s (string) , %d(integer) , %f(float) , you can specify to round off float by %.2f 
// you can use Sprintf to ensure the number is not sent to the standard output
// functions in go can have multiple return values in this case we wrap them in parenthesis
// functions in go will return named return values by default in order , known as a naked return
