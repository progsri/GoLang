package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	
	count := 0
	for i := 0; i < 100; i++ {
		go increment1(&count)
	}

	time.Sleep(time.Second * 5) // wait for the above goroutines to complete

	fmt.Println("Without locks ", count)  // this count will alwaysbe less 100 due to race conditions.

	channelRequests := make(chan int, 100)
	channelResponse := make(chan int, 100)

	go manageLock(channelRequests, channelResponse)
	

	for i := 0; i < 5; i++ {
		go increment2(channelRequests)
	}

	time.Sleep(time.Second * 50)  // wait for the above goroutines to complete
	countWithLock := <- channelResponse
	fmt.Println("With locks ", countWithLock)

	fmt.Println("waiting for input") 
	var input string
	fmt.Scanln(&input)

}

func increment1(count *int) {
	*count = *count + 1
}

func increment2(channel1 chan int) {
	channel1 <- 1
	fmt.Println("Sending 1 to channel")
	}

func manageLock(channel1 chan int, channel2 chan int){

	count := 0

	for {
		fmt.Println(" Blocking >> ")
		i := <- channel1
		fmt.Println("Recevied >> " + strconv.Itoa(i))
		if i > 0 {
			count = count + i
		}

		fmt.Println("sending to response channel count >> " + strconv.Itoa(count))
		channel2 <- count
	}
}
