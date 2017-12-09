package main

import (
	"fmt"
	"time"
)

func main() {
	
	count := 0
	for i := 0; i < 100; i++ {
		go increment1(&count)
	}

	time.Sleep(time.Second * 5) // wait for the above goroutines to complete

	fmt.Println("Without locks ", count)  // this count will alwaysbe less 100 due to race conditions.

	channelRequests := make(chan int, 100)
	channelResponse := make(chan int)

	go manageLock(channelRequests, channelResponse)
	
	countWithLock := 0
	for i := 0; i < 100; i++ {
		go increment2(channelRequests)
	}

	time.Sleep(time.Second * 5)  // wait for the above goroutines to complete
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
}

func manageLock(channel1 chan int, channel2 chan int){

	for i:= range channel2{

	}
}
