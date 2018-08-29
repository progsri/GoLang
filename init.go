package main

import (
	"fmt"
)

const singleConst = "single constants"

const (
	group1 = "g1"
	group2 = "g2"
)

var singleVar string = "single variable"

func main(){

	fmt.Println(singleConst)
	fmt.Println(singleVar)

	fmt.Println("MAIN")

	xPtr := new(int);
	*xPtr = 10;
	fmt.Println(&xPtr);
	fmt.Println(*xPtr);
}

func init(){
	// init always runs before main ..use to do initialization
	fmt.Println("INIT")
}

