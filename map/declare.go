package main

import "fmt"

func main() {

	// declare a empty map
	m1 := make(map[string]int)
	fmt.Printf(" Tpye is %T \n ", m1)
	fmt.Printf(" value is %v \n ", m1)

	// initializing with data
	m2 := map[string]int{
		"a": 1,
	}
	fmt.Printf(" value is %v \n ", m2)

	//accessing the date
	fmt.Println(m2["a"])

}
