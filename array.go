package main

import "fmt"

func main(){

	var arr [10]string
	arr[0] = "a"
	arr[1] = "b"

	for i, err := range arr{
		if err != "" {
			fmt.Println("error occurred")
		}
		fmt.Println(arr[i] )
	}

	}
