package main

import "fmt"

type  Formatter interface {
	Formart () string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (p PlainText) Formart () string {
	return fmt.Sprintf("%s", p.message)
}
func (b Bold) Formart () string {
	return fmt.Sprintf("**%s**",b.message)
}

func (c Code) Formart () string {
	return fmt.Sprintf("`%s`", c.message)
}

func sendMessage (formatter Formatter) string {
	return formatter.Formart()
}

func main() {
	plainText := PlainText{message: "Hello, world!"}
	boldText := Bold{message: "Hello, world!"}
	codeText := Code{message: "Hello, world!"}

	fmt.Println(sendMessage(plainText))
	fmt.Println(sendMessage(boldText))
	fmt.Println(sendMessage(codeText))
}

