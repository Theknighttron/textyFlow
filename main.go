package main

import "fmt"

// struct definition
type messageToSend  struct{
	PhoneNumber int
	Message string
}

func test(text messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", text.Message, text.PhoneNumber)
	fmt.Println("==================================")
}

func main() {
	test(messageToSend{
		PhoneNumber: +255621458592,
		Message: "Thanks for signing up",
	})
	test(messageToSend{
		PhoneNumber: +255741859302,
		Message: "Love to have you aboard",
	})
	test(messageToSend{
		PhoneNumber: +2557392869204,
		Message: "We're so excited to have you",
	})
}