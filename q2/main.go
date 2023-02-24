package main

import (
	"fmt"
	"time"
)

func trueToFalse(channel chan bool) {

	channel <- true

	fmt.Println("Message Sent")

}

func check(channel chan bool) {
	for !<-channel {
		fmt.Println("Awaiting Message")
	}

	fmt.Println("Message Recieved")

}
func main() {

	channel := make(chan bool, 1)

	channel <- false

	go check(channel)
	go trueToFalse(channel)

	time.Sleep(100 * time.Millisecond)

}
