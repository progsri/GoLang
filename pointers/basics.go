package main

import "fmt"

type container struct {
	id string
}

func main() {

	a := 1
	aPointer := &a
	fmt.Println(aPointer)

	//to dereference
	fmt.Println(*aPointer)

	b := container{id: "abc"}
	bPointer := &b
	fmt.Println(bPointer)

	//to deference ..with dot golang is doing something
	fmt.Println(bPointer.id)
	fmt.Println((*bPointer).id)
}
