package main

import (
	"fmt"
)

func main() {

	channel1 := make(chan string)

	// channel1 <- "main before method 1" // will throw error : all go routines are asleep ... since this is blocking call and waits to assign data with in a
	// goroutine as channels are only used via goroutines.
	go method1(channel1)

	fmt.Println(<-channel1) // we can always access the data from a channel.
}

func method1(channel1 chan string) {
	//channel1 <- "method1"
}
