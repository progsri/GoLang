package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	go sendTo(channel1)
	fmt.Println("blocking received >>", <-channel1) //blocking mode

	//non blocking mode  receiving
	go sendTo(channel1)
	select {
	case <-channel1:
		fmt.Println("non blocking >>  received")
	default:
		fmt.Println("non blocking >> did not receive")
	}

	//non blocking receive
	go sendToNonBlocking(channel1)

	fmt.Println("Wait for user Input")
	var input string
	fmt.Scanln(&input)
}

func sendTo(channel1 chan string) {
	time.Sleep(time.Second * 3)
	channel1 <- "sending to channel"
	fmt.Println("blocking >> sent to channel")
}

func sendToNonBlocking(channel1 chan string) {
	//There is ono receiver
	select {
	case channel1 <- "sending to non blocking channel":
		fmt.Println("non blocking sent")
	default:
		fmt.Println(" non blocking could not send")
	}
	fmt.Println("non blocking >> end")
}
