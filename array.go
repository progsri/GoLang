package main

import "fmt"

func main() {

	var arr [10]string
	arr[0] = "a"
	arr[1] = "b"

	for i, err := range arr {
		if err != "" {
			fmt.Println("error occurred")
		}
		fmt.Println(arr[i])
	}

	 arr1 := [...]string{"a","b","c"}
	 fmt.Println("before sending arr1 values ", arr1)
	 m1(arr1)
	 fmt.Println("after sending arr1 values ", arr1) 
	 // So Passing arraying is value based not reference based..unlike maps which is referenced based
}

func m1(arr [3]string) {
	arr[0] = "modfiying"
	fmt.Println(" arr is func ", arr)
}
