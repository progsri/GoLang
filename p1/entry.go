package main

import (
	"GoLang/p1/p2"
	"GoLang/p1/p2/p3"
	"fmt"
)

func main() {

	e1 := p3.Employee{Id: "e1", Email: "e1@gmail.com"}
	fmt.Println(e1)
	e1.Skills()
	//	fmt.Println(e1.Apply()) // This would not work as the Apply method is defined in a different package.

	s1 := p2.SubEmployee{e1}

	fmt.Println("Sub Employee ", s1.Apply())

}
