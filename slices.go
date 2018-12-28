package main

import "fmt"

func main() {

	// empty slice
	string_slice_5 := make([]string, 5)
	fmt.Println(string_slice_5)

	// Slice is a array
	string_slice_5[0] = "1"
	fmt.Println(string_slice_5)

	string_slice_5_copy := make([]string, len(string_slice_5))
	copy(string_slice_5_copy, string_slice_5)

	fmt.Println("original  ", string_slice_5)
	fmt.Println("copy ", string_slice_5_copy)

	//case 1
	// Changes to the sliced array effects the original array
	fmt.Println(" ------ case 1 #  Changes to the sliced array effects the original array --------")
	arr1 := [...]string{"a", "b", "c"}
	arr2 := arr1[1:3]
	fmt.Println("case1 before ", arr1)
	arr2[0] = "slice modified"
	fmt.Println("case1 after ", arr1)

	// Case 2
	// Changes to the original array effects the sliced array as well
	fmt.Println("Case2 # Changes to the original array effects the sliced array as well")
	arr3 := [...]string{"a", "b", "c"}
	arr4 := arr3[0:3]
	fmt.Println("case2 before ", arr4)
	arr3[0] = "slice modified"
	fmt.Println("case2 after ", arr4)

	// Capacity
	arr5 := [...]int{0,1,2,3,4,5,6}
	arr6 := arr5[2:3]
	fmt.Println("arr6 ", arr6)
	fmt.Println(" Length of arr6 ", len(arr6))
	fmt.Println(" Capacity of arr6 ", cap(arr6))

}
