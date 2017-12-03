package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	fmt.Println(reflect.TypeOf(ticker))

	go tickEveryOneSecond(ticker)

	time.Sleep(time.Second * 10)
	ticker.Stop() // stop the ticking

	fmt.Println("waiting for user input")
	var input string
	fmt.Scanln(&input)
}

func tickEveryOneSecond(ticker *time.Ticker) {

	for tick := range ticker.C {
		fmt.Println(" Ticked at >> ", tick)
	}
}
