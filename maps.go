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
	fmt.Println(" When there is no key the first value will return the default value ( zero-valued) \n 			So, always check for the secon value... which tells about the existence.")

	map_1["key3"] = 111
	v1, v2 := map_1["key3"]

	fmt.Println("When key exists >> first is the value of key >> ", v1)
	fmt.Println("When key exists >> second is the existence of the key >> ", v2)

	v3, v4 := map_1["key4"]

	fmt.Println("When key exists >> first is the value of key >> ", v3)
	fmt.Println("When key exists >> second is the existence of the key >> ", v4)

	fmt.Println("---------- Another way of creating a map -------")

	map_2 := map[string]int{}
	map_2["key1"] = 1

	fmt.Println(" map 2 ", map_2)

	// Is passing map reference based or value based
	fmt.Println("-----  passing map reference based unlike arrays which is value based --------")
	map1 := map[string]int {
		"first" : 1,
		"second" : 2,
	}
	fmt.Println("before passing to a func ", map1);
	m1(map1)
	fmt.Println("after passing to a func ", map1);


}

func m1(map2 map[string]int) {
	map2["first"] = 3
}
