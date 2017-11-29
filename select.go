package main

import (
	"fmt"
	"time"
)

func main() {

	channelServer1 := make(chan string)
	channelServer2 := make(chan string)

	fmt.Println("-------- with out select --------")
	// without select

	nonSelectStart := time.Now()
	go server_1(channelServer1)
	go server_2(channelServer2)

	fmt.Println("Blocking call 1 ", <-channelServer1) // Blocking call
	fmt.Println("Blocking call 1 ", <-channelServer2) // Blocking call

	// Waits till the above 2 blocking call are done.

	fmt.Println(" executes after the above 2 blocking calls respond")
	//	nonSelectStop := time.Now().Nanosecond
	nonSelectTimeTaken := time.Since(nonSelectStart)
	fmt.Println("Time taken for non select ", nonSelectTimeTaken)

	selectStart := time.Now()
	go server_1(channelServer1)
	go server_2(channelServer2)

	select { // waits only till one blocking call is done
	case m1 := <-channelServer1:
		fmt.Println("select >> ", m1)
	case m2 := <-channelServer2:
		fmt.Println("select >> ", m2)
	}

	selectTimeTaken := time.Since(selectStart)
	fmt.Println("Time taken for  select ", selectTimeTaken)

	var input string
	fmt.Scanln(&input)
}

func server_1(channel1 chan string) {
	time.Sleep(10000)
	channel1 <- "server responds 1" // Responds in 10 secs
	fmt.Println("server 1 responded")
}

func server_2(channel2 chan string) {
	time.Sleep(5000)
	channel2 <- "server responds 2" // Responds in 5 secs
	fmt.Println("server 2 responded")
}
