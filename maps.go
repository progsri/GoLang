package main

import "fmt"

func main() {

	fmt.Println("-------creating map-------------")
	map_1 := make(map[string]int)

	map_1["key1"] = 1

	fmt.Println("adding first key to map >> ", map_1)

	map_1["key2"] = 2

	fmt.Println("adding a second key to map >> ", map_1)

	delete(map_1, "key2")

	fmt.Println("deleting key from a map >> ", map_1)

        fmt.Println("------- map's key returns multi values ---------------")

	map_1["key3"] = 0
	v1, v2 := map_1["key4"]

	fmt.Println("first is the value of key >> ", v1)
	fmt.Println("second is the existence of the key >> ", v2)
}
