package main

import "fmt"

type Person struct {
	name string
	zip  int
}

type PersonAnother struct {
	name string
	zip  int
}

func main() {
	p1 := Person{
		"first", 111,
	}

	p2 := Person{
		"first", 111,
	}

	if p1 == p2 {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("Not Equal")
	}

	p3 := Person{
		"first", 222,
	}

	if p1 == p3 {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("Not Equal")
	}

	// p4 := PersonAnother{
	// 	"first", 111,
	// }

	//if p1 == p4 {   //wow :: Go does not even let you to compile
	//	fmt.Println("EQUAL")
	//} else {
	//	fmt.Println("Not Equal")
	//}

}
