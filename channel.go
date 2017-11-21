package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("------ creating a channel -----------")
	fmt.Println(" Channels should be used only with goroutine ")

	commonChannel1 := make(chan string)
	commonChannel2 := make(chan string)
	commonChannel3 := make(chan string)

	// after 10 secs commonChannel1 has the required value
	go method1(commonChannel1)

	// only when the above commonChannel1 puts the required value commonChannel2 will have the required value
	go method2(commonChannel1, commonChannel2)

	// only when the above commonChannel2 puts the required value commonChannel3 will have the required value
	go method3(commonChannel2, commonChannel3)

	//Wait till the go rountines complete
	var input string
	fmt.Scanln(&input)

}

func method1(channel1 chan<- string) {
	time.Sleep(100000000)
	channel1 <- "method_1_proceed" // copy to this channel

	fmt.Println("--- end of method 1 --- ")
}

func method2(channel2 <-chan string, channel3 chan<- string) {

	valueFromChannel2 := <-channel2

	fmt.Println("channel2 has a value ", valueFromChannel2)
	for strings.Compare(valueFromChannel2, "method_1_proceed") == 0 {
		time.Sleep(10000)
		channel3 <- "method_2_proceed"
		fmt.Println("assigned value to channel3")
		//break
	}

	fmt.Println("--- end of method 2 --- ")
}

func method3(channel4 <-chan string, channel5 chan<- string) {

	valueFromChannel4 := <-channel4
	for strings.Compare(valueFromChannel4, "method_2_proceed") == 0 {
		time.Sleep(10000)
		channel5 <- "method_3_proceed"
		fmt.Println("assigned value to channel5")
		break
	}

	fmt.Println("--- end of method 3 --- ")
}
