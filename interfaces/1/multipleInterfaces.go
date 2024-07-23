package main

import "fmt"

type expense interface {
	cost() float64
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}


func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}else{
		return float64(len(e.body)) * 0.02
	}
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}else {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
}



func test (e expense , f formatter) {
	fmt.Println(e.cost())
	fmt.Println(f.format())
}

// you can name the variables in the function definition of your interface 
type Copier interface {
	Copy (sourceFile , destinationFile string) (bytesCopied int)
}

func main () {
	email1 := email{
		isSubscribed: true,
		body: "Hello World",
	}

	email2 := email{
		isSubscribed: false,
		body: "How are you",
	}

	test(email1, email2)
}