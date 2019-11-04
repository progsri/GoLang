package main

import "fmt"
import "time"

func parallel(channel chan string) {
	fmt.Println("at parllel")
	time.Sleep(10000 * time.Millisecond)
	channel <- "loaded"
	//close(channel)
}

func main() {
	fmt.Println(" start ")
	channel := make(chan string, 1)
	channel <- "1"
	channel <- "2"
	go parallel(channel)

	value := <-channel

	fmt.Println(" end " + value)

}
