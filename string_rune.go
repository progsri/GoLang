package main

import "fmt"

func main() {

	fmt.Println("--------- immuatbale strings ")
	s := "strings" // immutable strings
	fmt.Println(s)
	// These are stored a unicode ...encode a UTF-8

	fmt.Println(s[0])         // in uTF-8 encoded
	fmt.Printf("%c \n", s[0]) // in human readable format

	fmt.Println("----- Making Strings mutable -----")

	r1 := []rune(s) // gives a slice of type rune
	fmt.Println(r1)
	fmt.Println(append(r1, 'n'))
	r1[0] = 't'
	fmt.Println(s)
	fmt.Println(string(r1))

}
