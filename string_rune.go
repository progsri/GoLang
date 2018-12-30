package main

import "fmt"

func main() {
	s := "strings" // immutable strings
	// These are stored a unicode ...encode a UTF-8

	fmt.Println(s[0])         // in uTF-8 encoded
	fmt.Printf("%c \n", s[0]) // in human readable format
}
