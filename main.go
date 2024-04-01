package main

import (
	"fmt"
)

// Using nested struct

type messageToSend struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func test(mToSend messageToSend) {
	fmt.Printf("sending '%s' from %s \n", mToSend.message, mToSend.sender.name)
	fmt.Printf("%s (%v) to %s (%v)...\n", mToSend.sender.name, mToSend.sender.number, mToSend.recipient.name, mToSend.recipient.number)
	fmt.Println("...sent! ")
	fmt.Println("=======================================")
}


func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{name: "Carter", number: +255739683865},
		recipient: user{name: "Tailor Jin", number: +255728694853},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{name: "Suzan", number: +255784926519},
		recipient: user{name: "James Peter", number: +255793029346},
	})
	test(messageToSend{
		message: "you have a pary tommorow",
		sender: user{number: +255638295723},
		recipient: user{name: "Sally Sue"},
	})
}
