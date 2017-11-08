package main

import "fmt"

func main(){
	takes_var_args("a", "b", "c")
	
	// slice
	slice := make([]string,3)
	slice[0] = "1"
	slice[1] = "0"
	takes_var_args(slice...) //  slice should have ... if the calling 
				// method takes variable arguments.
}

func takes_var_args(arr ...string){  // needs ... to signify variable arguments
	for _, v := range arr{
	   fmt.Println(v)
	}
}
