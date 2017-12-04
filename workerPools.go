package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go worker(channel1, channel2)
	go worker(channel1, channel2)
	go worker(channel1, channel2)

	channel1 <- "a"
	fmt.Println("[Assing] job assigned ")

	channel1 <- "b"
	fmt.Println("[Assing] job assigned ")

	channel1 <- "c"
	fmt.Println("[Assing] job assigned ")

	// channel1 <- "d"
	// fmt.Println("[Assing] job assigned ")

	close(channel1)

	for result := range channel2 {
		fmt.Println("[Assign] result >> " + result)
	}

	fmt.Println("waiting for user input")
	var input string
	fmt.Scanln(&input)

}

func worker(recvChannel <-chan string, sendChannel chan<- string) {

	for {
		fmt.Println("[worker] >> lookin for job")
		job := <-recvChannel
		fmt.Println("[worker] >> job received " + job)
		time.Sleep(time.Second * 2)
		sendChannel <- job + " " + time.Now().String()
		fmt.Println("[worker] >> job result posted ")
	}
}
