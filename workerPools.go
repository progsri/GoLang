package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	channel1 := make(chan string, 20)
	channel2 := make(chan string)

	go worker(channel1, channel2)
	go worker(channel1, channel2)
	go worker(channel1, channel2)

	for i := 0; i < 10; i++ {
		channel1 <- "a" + strconv.Itoa(i)
	}

	channel1 <- "d"

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
		//fmt.Println("[worker] >> lookin for job")
		job := <-recvChannel // This is not a blocking call ...it something exists it will get otherwise will be empty.

		if len(job) > 0 {
			fmt.Println("[worker] >> job received " + job)
			time.Sleep(time.Second * 2)
			sendChannel <- job + " " + time.Now().String()
			fmt.Println("[worker] >> job result posted ")
		}
	}
}
