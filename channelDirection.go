package main

import (
	"fmt"
)

func main() {

	fmt.Println("----- channel direction -----")

	channel1 := make(chan string)
	go method3(channel1)
	fmt.Println(<-channel1)

	go method4(channel1)

	var input string
	fmt.Scanln(&input)

}

func method3(channel1 chan<- string) {
	fmt.Println("Trying to send to this channel")
	channel1 <- "send to method 3" //there is a channel  to receive this message simultaneously ...so though it is a blocking call
	// looks like a non blocking call.
	fmt.Println("sent to the channel")
}

func method4(channel1 chan<- string) {
	fmt.Println("Trying to send to this channel")
	channel1 <- "send to method 4" //as there is no channel to receive simultaneously so this will be blocked
	fmt.Println("sent to the channel")
}
