package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch p := n.(type) {
	case directMessage:
		return p.senderUsername, p.importance()
	case groupMessage:
		return p.groupName, p.importance()
	case systemAlert:
		return p.alertCode, p.importance()
	default:
		return "", 0
	}
}

func main() {
	dm := directMessage{
		senderUsername: "user123",
		messageContent: "Hey, check this out!",
		priorityLevel:  10,
		isUrgent:       true,
	}

	gm := groupMessage{
		groupName:      "Developers",
		messageContent: "Meeting at 10 AM",
		priorityLevel:  5,
	}

	sa := systemAlert{
		alertCode:      "ALERT001",
		messageContent: "System maintenance scheduled",
	}

	notifications := []notification{dm, gm, sa}

	for _, n := range notifications {
		info, importance := processNotification(n)
		fmt.Printf("Notification from %s with importance %d\n", info, importance)
	}
}
