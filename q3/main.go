package main

import (
	"fmt"
	"time"
)

func processOne() {

	fmt.Println("Process 1 completed")

}

func processTwo() {
	fmt.Println("Process 2 completed")

}
func main() {
	go processOne()
	go processTwo()

	fmt.Println("process 3")
	time.Sleep(1000 * time.Millisecond)

}
