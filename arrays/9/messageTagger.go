package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for _, message := range messages {
		message.tags = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}

	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Sales")
	}
	return tags
}


func main () {
	messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
	taggedMessages := tagMessages(messages, tagger)
	fmt.Println(taggedMessages)
}