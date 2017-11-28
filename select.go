package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("-------- with out select --------")
	// without select
	channelServer1 := make(chan string)
	channelServer2 := make(chan string)

	go server_1(channelServer1)
	go server_2(channelServer2)

	<-channelServer1 // Blocking call
	<-channelServer2 // Blocking call

	fmt.Println(" executes after the above 2 blocking calls respond")

}

func server_1(channel1 chan string) {
	time.Sleep(10000)
	channel1 <- "server responds" // Responds in 10 secs
	fmt.Println("server 1 responded")
}

func server_2(channel2 chan string) {
	time.Sleep(5000)              // Responds in 5 secs
	channel2 <- "server responds" // Responds in 10 secs
	fmt.Println("server 2 responded")
}
