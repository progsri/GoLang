package main

import "fmt"

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go method1(channel1)
	go method2(channel2)
	go method3(channel1, channel2)

	var input string
	fmt.Scanln(&input) // Why does scan need a address ?

}

func method1(channel1 chan string) {
	//	channel1 <- "write to channel1"
}

func method2(channel2 chan string) {
	//	channel2 <- "write to channel2"
}

func method3(channel1 chan string, channel2 chan string) {
	a, b := <-channel1, <-channel2

	fmt.Println("a is ", a, " :: b is ", b)
}
