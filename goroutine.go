package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("------------ reg --------------")

	regStart := time.Now().Nanosecond()
	regularMethod("regular") // Synchronous
	regStop := time.Now().Nanosecond()
	fmt.Println(regStop - regStart)

	fmt.Println("------------ goroutine --------------")

	goStart := time.Now().Nanosecond()
	go regularMethod("go routine 1")  // Asynchronous
	go regularMethod("go routine 2")  // Asynchronous
	go regularMethod("go routine 3")  // Asynchronous
	go regularMethod("go routine 4")  // Asynchronous
	go regularMethod("go routine 5")  // Asynchronous
	go regularMethod("go routine 6")  // Asynchronous
	go regularMethod("go routine 7")  // Asynchronous
	go regularMethod("go routine 8")  // Asynchronous
	go regularMethod("go routine 9")  // Asynchronous
	go regularMethod("go routine 10") // Asynchronous
	goStop := time.Now().Nanosecond()
	fmt.Println(goStop - goStart)

	// to wait for the go rountine to finish execution
	var value string
	fmt.Scanln(&value)

	fmt.Println("------------ end --------------")
}

func regularMethod(caller string) {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %s - %d \n ", caller, i)
	}
}
