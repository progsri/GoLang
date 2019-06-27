package main

import "errors"
import "fmt"

func main() {

	// Lets you create a errorString struct
	s1 := errors.New("error 1")
	fmt.Println(s1)

	// Gives you access to the struct method Error()
	fmt.Println(s1.Error())
}

