package main

import (
	"GoLang/p1/p2"
	"GoLang/p1/p2/p3"
	"fmt"
)

// NewEmployee .... ( this is for go-lint)
type NewEmployee interface {
	Vacation()
}

func listOfVacationPlans(p NewEmployee) {
	fmt.Printf("Type = %T, value = %v\n", p, p)
}

// Universal .. this has no methods ..so all structs can be of this type
type Universal interface{}

func listOfSomething(p Universal) {
	fmt.Printf("Type = %T, value = %v\n", p, p)
}

func main() {

	e1 := p3.Employee{Id: "e1", Email: "e1@gmail.com"}
	fmt.Println(e1)
	e1.Skills()
	//	fmt.Println(e1.Apply()) // This would not work as the Apply method is defined in a different package.

	s1 := p2.SubEmployee{e1}

	fmt.Println("Sub Employee ", s1.Apply()) // This calls SubEmployee

	listOfVacationPlans(e1) // works since p3.Employee implements NewEmployee interface
	listOfSomething(e1)     // even though p3.Employee does not implement Universal interface, this still works
	// as Universal Interface is Empty Interface ( no methods )
}
