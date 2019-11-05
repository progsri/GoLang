package main

import "fmt"

type s1 struct {
	name string
	s2
}

type s2 struct {
	id string
}

func (s s1) m1() {
	fmt.Println("called m1")
}

func (s s2) m2() {
	fmt.Println("called m2")
}

func main() {
	v1 := s1{name: "s1 struct"}
	fmt.Printf(" v1 %v ", v1)
	v1.m1()
	v1.m2()
}
