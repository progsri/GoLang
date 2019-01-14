package main

import (
	"fmt"
)

func main() {

	channel1 := make(chan string)

	// channel1 <- "main before method 1" // will throw error : all go routines are asleep ... 
	go method1(channel1)
	go method2(channel1)
	fmt.Println("continue with main go routine")

	// all go routines are asleep ... deadlock
	// if a channel in main go routine is waiting to read/write , but there is no write/read happening, then
	// i think main go routine knows that and throws all go routines are asleep ... deadlock


	//We need the below as
	// as main go routine does not wait for other go routines to complete
	var input string
	fmt.Println("Scan will make sure ...mian go rountine does not die.")
	fmt.Scanln(&input)
}

func method1(channel1 chan string) {
	fmt.Println("Start writing to channel")
	//channel1 <- "method1"
	fmt.Println("Done writing to channel")
}

func method2(channel1 chan string) {
	fmt.Println("Start reading to channel")
	fmt.Println(<-channel1) // we can always access the data from a channel...if someone is sending to it
	fmt.Println("Done reading to channel")
}
