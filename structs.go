package main

import (
	"fmt"
	"reflect"
)

type group struct {
	name string
	count int
	sub subgroup
}

type subgroup struct {
	id string
	desc string
}

func main() {

	fmt.Println("-------------creating a struct ---------")
	subg1 := subgroup{"s1","sbg1"}
	fmt.Println(subg1)
	g1 := group{"g1", 1, subg1}
	fmt.Println(g1)

	g2 := &group{
		name : "g1",
		count : 1,
		sub : subg1}
	fmt.Println("g2 " , g2)
	fmt.Println("type of g2 ", reflect.ValueOf(g2).Kind())
	fmt.Println("type of &g2 ", &g2)

	fmt.Println("---------access struct ----------------")
	fmt.Println("directly >>> ", g1.name)
	// if name should be accessed outside this file then the first letter should be capital.

	g1_ptr := &g1
	fmt.Println("pointer >> ", g1_ptr.name)

	g1_ptr.sub.id = "test"
	fmt.Println("update >>", g1)
}

