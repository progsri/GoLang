package main

import "fmt"

type group struct {
	name string
	count int
	subgroup subgroup
}

type subgroup struct {
	id string
	desc string
}

func main() {
	
	subg1 := subgroup{"s1","sbg1"}
	fmt.Println(subg1)
	g1 := group{"g1", 1, subg1}
	fmt.Println(g1)
	
}

