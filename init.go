package main

import (
	"fmt"
)

func main(){
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