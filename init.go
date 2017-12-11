package main

import (
	"fmt"
)

func main(){
	fmt.Println("MAIN")
}

func init(){
	// init always runs before main ..use to do initialization
	fmt.Println("INIT")
}