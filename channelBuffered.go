package main

import (
	"fmt"
)

func main() {

	channel1 := make(chan string, 2) //buffered channel

	go method2(channel1)

	fmt.Println(<-channel1)

	fmt.Println("--- end---")

	var input string
	fmt.Scanln(&input)

	fmt.Println(<-channel1)
}

func method2(channel1 chan string) {
	channel1 <- "method2 1" // there should be a receiver at the very moment for us to send to this channel ...
	// otherewise it will be blocked.
	fmt.Println("sent to channel >> msg >> method2 1")

	channel1 <- "method2 2" // Since this is buffered channel inspite of no one is receiving it at that moment
	//... the message is sent thru
	fmt.Println("sent to channel >> msg >> method2 2")
}
