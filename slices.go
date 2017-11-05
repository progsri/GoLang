package main

import "fmt"

func main() {

	// empty slice
	string_slice_5 := make([]string, 5);
	fmt.Println(string_slice_5)

	// Slice is a array
	string_slice_5[0] = "1"
	fmt.Println(string_slice_5)

	string_slice_5_copy := make([]string, len(string_slice_5))
	copy(string_slice_5_copy, string_slice_5)

	fmt.Println("original  " ,  string_slice_5)
	fmt.Println("copy ", string_slice_5_copy)

}
