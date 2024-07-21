package main

import "fmt"

type membershipPlan string

const (
	Standard membershipPlan = "standard"
	Premium membershipPlan = "premium"
)

type Membership struct {
	Type membershipPlan 
	MessageCharLimit int
}

type User struct {
	Name string 
	Membership
}

func NewUser (name string , p membershipPlan) User {
	limit := 0
	if p == Standard {
		limit = 100

	}else {
		limit = 1000
	}
	return User {
		Name: "Richard",
		Membership: Membership {
			Type: p,
			MessageCharLimit: limit,
		},
	}
}

func main () {
	user1 := NewUser ("Richard" , Standard)
	fmt.Printf("%v \n", user1)
}


