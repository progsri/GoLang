package main

import "fmt"

type config struct {
	addressNop i1
	addressP   i1
	network    string
}

func main() {

	structNop := S1{
		"s1-nop",
	}
	fmt.Println(structNop)

	structP := S1{
		"s1-p",
	}
	fmt.Println(structP)

	structConfig := config{
		addressNop: structNop,
		addressP:   structP,
		network: "tcp",
	}

	fmt.Println(structConfig)

}
