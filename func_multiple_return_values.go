package main

import "fmt"

func main(){

	fmt.Println("--------- No return value -------- ")
	
	fmt.Println(return_no_value)

	fmt.Println("-------- one return value --------" )

	fmt.Println(return_one_value("only one value"))
}

func return_no_value(value string) {
}

func return_one_value(value string) (string) {
	return value
}
