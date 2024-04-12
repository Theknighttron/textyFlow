package main

import (
	"fmt"
	"time"
)


// Interface
type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

type sendingReport struct {
	reportName string
	noOfSends  int
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s \n", bm.recipientName, bm.birthdayTime)
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is ready. You've sent %v messages \n", sr.reportName, sr.noOfSends)
}

func sendMessage(msg message) {
	fmt.Printf(msg.getMessage())
}

func test(m message) {
	sendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	test(sendingReport{
		reportName: "First Report",
		noOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Carter James",
		birthdayTime: time.Date(2000, 12,25, 0,0,0,0,time.UTC),
	})

}
