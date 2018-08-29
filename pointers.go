package main

import "fmt"
import "reflect"

type group1 struct {
	name string
	count int
	sub subgroup1
}

type subgroup1 struct {
	id string
	desc string
}

func main() {

	//  variable of type int holds int value
	i := 3
	fmt.Println("type of i ", reflect.ValueOf(i).Kind())
	fmt.Println("value of i >> ", i )
    fmt.Println("address of i >> ", &i)

	i_ptr  := &i
	*i_ptr = 4
	fmt.Println("type of i_ptr ", reflect.ValueOf(i_ptr).Kind())
	fmt.Println("pointer(address) of i >> ", i_ptr)
	fmt.Println("pointer(address) of i >> ", &i_ptr) //?? What is this actually giving
	fmt.Println("value of pointer >> ", *i_ptr)
	
	*i_ptr = 4
	fmt.Println("-----------------------------------")
	fmt.Println("type of i_ptr ", reflect.ValueOf(i_ptr).Kind())
	fmt.Println("changing the value of the address >> ", i)

	fmt.Println("value of i >> ", i )
    fmt.Println("address of i >> ", &i)
    fmt.Println("pointer(address) of i >> ", i_ptr)
	fmt.Println("value of pointer >> ", *i_ptr)

	fmt.Println("--------------------")

	xPtr := new(int)
	*xPtr = 10
	fmt.Println(xPtr)
	fmt.Println(&xPtr) //?? What is this actually giving
	fmt.Println(*xPtr)

	fmt.Println("--------------------")

	subg1 := subgroup1{"s1","sbg1"}
	g1 := group1{
		name : "g1",
		count : 1,
		sub : subg1}
	fmt.Println("g1 " , g1)
	fmt.Println("type of g1 ", reflect.ValueOf(g1).Kind())
	fmt.Println("type of &g1 ", &g1)

	fmt.Println("--------------------")

	subg2 := subgroup1{"s1","sbg1"}
	g2 := &group1{
		name : "g1",
		count : 1,
		sub : subg2}
	fmt.Println("g2 " , g2)
	fmt.Println("type of g2 ", reflect.ValueOf(g2).Kind())
	fmt.Println("type of &g2 ", &g2)

	fmt.Println("--------------------")

}
