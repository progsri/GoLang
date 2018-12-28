package main

import "fmt"

func main() {

	// empty slice
	string_slice_5 := make([]string, 5)
	fmt.Println(string_slice_5)

	// Slice is a array
	string_slice_5[0] = "1"
	fmt.Println(string_slice_5)

	fmt.Println("-----------Copying a slice does not modify the original array which can be garbage collected-----------------")
	string_slice_5_copy := make([]string, len(string_slice_5)) // Creates a new slice which can be used to copy
	copy(string_slice_5_copy, string_slice_5)

	fmt.Println("original  ", string_slice_5)
	fmt.Println("copy ", string_slice_5_copy)
	string_slice_5_copy[0] = "modified"
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
	fmt.Println("----- Capacity -----")
	arr5 := [...]int{0, 1, 2, 3, 4, 5, 6}
	arr6 := arr5[2:3]
	fmt.Println("arr6 ", arr6)
	fmt.Println(" Length of arr6 ", len(arr6))
	fmt.Println(" Capacity of arr6 ", cap(arr6))

	//Append
	fmt.Println("----- Append ------")
	arr7 := [...]int{0, 1, 2, 3, 4, 5, 6}
	arr8 := arr7[2:4]
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(" Length of arr8 ", len(arr8), " Capacity of arr8 ", cap(arr8))
	arr9 := append(arr8, 7) // append creates a new slice and this slice is attached to the original array itself
	fmt.Println("\t after append to slice arr8")
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	arr7[2] = 999
	fmt.Println(" \t after modifying the original array arr7")
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	arr8[1] = 888
	fmt.Println(" \t after modifying the first slice arr8")
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	arr10 := append(arr9, 1111)
	fmt.Println(arr10)
	fmt.Println(" append creates a new slice and this slice is attached to the original array itself")
}
