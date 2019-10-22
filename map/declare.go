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

	// copying a map to another map still references the original map....
	// to not hae a reference use for loop
	m3 := m2
	fmt.Printf(" value of m3 is %v \n ", m3)
	delete(m2, "a")
	fmt.Printf(" value of m3 is %v \n ", m3)

}
