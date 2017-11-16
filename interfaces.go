package main

import "fmt"

func main() {
	
	first_str := first_struct{}
	fmt.Println(" calling the struct method ( not interface ) >> ", first_str.method() )

	second_str := second_struct{}
	fmt.Println(" calling the struct method ( not interface ) >> ", second_str.method() )

	//Interface power
	power_of_interface(first_str)

	power_of_interface(second_str)
}

type  I interface{
	method() int
}

type  first_struct struct{
}

// thru struct implementing interface I method ( indirectly means you are implementing the interface methods ..so this struct is of type I )
// Do not use func(s *first_struct) method() int i.e struct pointer
func (s first_struct) method() int{
	return 1
}

type second_struct struct{
}

func (o second_struct) method() int {
	return 2;
}


func  power_of_interface(i I){
	fmt.Println(i.method())
}

