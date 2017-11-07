package main

import "fmt"

func main(){
	
	fmt.Println("-------------- slices -----------------")
	slice := make([]string, 6)
	slice[0] = "a"
	slice[2] = "b"
	slice[5] = "c"

	fmt.Println(" slice " , slice)
	
	for v1, v2 := range slice {
		fmt.Println("First Value [index] >> ", v1)
		fmt.Println("Second Value [value] >> ", v2)
	}

	fmt.Println("------------------string -----------------")

	var word string = "this is a sentence"
	for v3, v4 := range word{
		fmt.Println("First Value [index] >> ", v3)
		fmt.Println("Second Value [value] >> ", v4)
	}

	fmt.Println("-------------------- map -----------------")
	
	map_1 := map[string]string{}
	map_1["k1"] = "v1"
	map_1["k2"] = "v2"
	map_1["k3"] = "v3"

	for k,v := range map_1 {
		fmt.Printf("key %s >> value %s \n",k, v) 
	}
}
