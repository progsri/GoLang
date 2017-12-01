package main

import (
	"fmt"
)

func main() {

	channel1 := make(chan string, 1)

	channel1 <- "from main sent to channel1"
	close(channel1)

	//channel1 <- "after close send" // Will stop the execution by throwing feature/20171206_inHouseMsa_Itr3_Announcement
	fmt.Println("recv 1 ", <-channel1)
	fmt.Println("recv 2 ", <-channel1) // Does not throw any error..if something exists in the buffer will retreive it

	fmt.Println("Waiting for input")
	var input string
	fmt.Scanln(&input)
}
