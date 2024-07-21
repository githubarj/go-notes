package main 

import "fmt"

type messageToSend struct {
	message string
	sender user
	recipient user
}

func test(m messageToSend) {
	fmt.Printf("Sending message: %s to %d\n", m.message, m.sender.number)
}

type user struct {
	name string
	number int
}

func canSendMessage (mToSend messageToSend) bool {
	if mToSend.recipient.name != "" && mToSend.recipient.number != 0{
		if mToSend.sender.name != "" && mToSend.sender.number != 0 {
			return true
		}else {
			return false
		}
	}
	return false
}
type car struct {
	Make string
	Model string
	Height int
	Width int
	Wheel struct {
		Radius int
		Material string
	}
}
// go had embedded structs where you kind of extend another struct
type truck struct {
	Type string
	car
}

type rect struct {
	width int 
	height int
}
// we can declare methods for structs
func (r rect) area() int {
	return r.width * r.height
}

type authenticationInfo struct {
	password string
	username string 
}

func (auth authenticationInfo ) getBasicAuth () string {
	return fmt.Sprintf("Authentication: Basic %s:%s ",auth.username, auth.password) 
}

func main(){
	test(messageToSend{
		sender: user {
			name:"John",
			number: 254712345678,
		},
		message: "I am learning go",
	})
	fmt.Println(canSendMessage(messageToSend {
		sender: user {
			name: "John",
			number: 24453647152,
		},
		message: "hello",
		recipient: user {
			name: "Doe",
			number: 942894979498,
		},
		}))
		myCar := car{}
		fmt.Println(myCar)
		r := rect {
			width: 5,
			height: 4,
		}
		fmt.Println(r.area())
		user := authenticationInfo {
			username: "jonndoe",
			password:  "password123",
		}
		fmt.Println(user.getBasicAuth())
		// when declaring or initializing an embedded struct 
		volvo := truck {
			Type : "Heavy Duty",
			car : car {
				Make : "volvo",
				Model: "xc90",
				Height: 2,
				Width: 5,
				Wheel: struct {
					Radius int
					Material string
				}{ 
					Radius : 5,
					Material: "Rubber",
				},
			},
		}
		// when accessing the values we access them as top level keys 
		fmt.Printf("The truck details are type: %s , make: %s , wheel radius : %d \n", volvo.Type, volvo.Make , volvo.Wheel.Radius)

	}

	// the way you order data in a struct should be from largest to smallest as it has a big impact on memory 
