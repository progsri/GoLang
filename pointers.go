package main

import "fmt"

func main() {

	//  variable of type int holds int value
	i := 3
	fmt.Println("value of i >> ", i )

	fmt.Println("address of i >> ", &i)

	// variable of type *int holds address of the variable
	i_ptr  := &i

	fmt.Println("pointer(address) of i >> ", i_ptr)

        fmt.Println("value of pointer >> ", *i_ptr)

}
