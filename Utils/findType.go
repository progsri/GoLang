package main

import "fmt"

type test struct {
	id string
}

func main() {
	v := test{id: "v1"}
	fmt.Printf("value %v ", v)
	fmt.Printf("value %v ", v.(Type))
}
