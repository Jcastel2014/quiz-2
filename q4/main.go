package main

import (
	"fmt"
	"sync"
)

func trueToFalse(channel chan bool, waitGroup *sync.WaitGroup) {

	channel <- true

	fmt.Println("Message Sent")

	waitGroup.Done()

}

func check(channel chan bool, waitGroup *sync.WaitGroup) {
	for !<-channel {
		fmt.Println("Awaiting Message")
	}

	fmt.Println("Message Recieved")

	waitGroup.Done()

}
func main() {

	channel := make(chan bool, 1)

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	channel <- false

	go check(channel, &waitGroup)
	go trueToFalse(channel, &waitGroup)

	waitGroup.Wait()

}
