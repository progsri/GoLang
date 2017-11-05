package main

import "fmt"

func main() {

	fmt.Println("-------- create a map -------------")
	map_1 := make(map[string]int)

	map_1["con"] = 1

	fmt.Println("map ", map_1)

	map_1["key2"] = 2

	fmt.Println("map ", map_1)

	delete(map_1, "key2")

	fmt.Println("map ", map_1)


}
