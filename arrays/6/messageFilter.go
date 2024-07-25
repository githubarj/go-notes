package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct { // implements the message interface
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	filtered := []Message {}
	for _ , message := range messages {
		if message.Type() == filterType {
			filtered = append(filtered, message)
		}	
	}

	return filtered
}


// Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

// Assignment
// Your task is to implement a function that filters a slice of messages based on the message type.

// Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

func main() {
    messages := []Message{
        TextMessage{Sender: "Alice", Content: "Hello, World!"},
        MediaMessage{Sender: "Bob", MediaType: "image", Content: "photo.jpg"},
        LinkMessage{Sender: "Charlie", URL: "https://example.com", Content: "Check this out!"},
        TextMessage{Sender: "Dave", Content: "Another message"},
    }

    // Filter by "text"
    filteredByText := filterMessages(messages, "text")
    fmt.Println("Filtered by 'text':")
    for _, msg := range filteredByText {
        fmt.Println(msg)
    }

    // Filter by "media"
    filteredByMedia := filterMessages(messages, "media")
    fmt.Println("\nFiltered by 'media':")
    for _, msg := range filteredByMedia {
        fmt.Println(msg)
    }

    // Filter by "link"
    filteredByLink := filterMessages(messages, "link")
    fmt.Println("\nFiltered by 'link':")
    for _, msg := range filteredByLink {
        fmt.Println(msg)
    }
}