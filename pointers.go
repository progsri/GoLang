package main

import "fmt"

func main() {

	//  variable of type int holds int value
	i := 3

	fmt.Println("value of i >> ", i )
    fmt.Println("address of i >> ", &i)
	i_ptr  := &i
	fmt.Println("pointer(address) of i >> ", i_ptr)
	fmt.Println("value of pointer >> ", *i_ptr)
	
	*i_ptr = 4
	fmt.Println("-----------------------------------")
	fmt.Println("changing the value of the address >> ", i)

	fmt.Println("value of i >> ", i )
    fmt.Println("address of i >> ", &i)
    fmt.Println("pointer(address) of i >> ", i_ptr)
	fmt.Println("value of pointer >> ", *i_ptr)


}
