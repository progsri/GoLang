package main

import "fmt"

func main() {

	fmt.Println("--------- No return value -------- ")

	//fmt.Println(return_no_value("no value"))  // Since this does not return a value and Prinln requires
	// a value, this will throw compilation errors.

	fmt.Println("-------- one return value --------")

	fmt.Println(return_one_value("only one value"))

	fmt.Println("------- multiple return values ------ ")

	fmt.Println(return_multiple_values())

	fmt.Println("----- blank identifier ---------")

	_, _, v3 := return_multiple_values() // _ blank identifier is the place holder to ignore the values and yet helps to maintain the return values order.
	fmt.Println(v3)

	var f1, e1, _ string = return_multiple_values_with_namedReturn() // Using Blank identifier to discard
	fmt.Println( e1, f1)
}

func return_no_value(value string) {
}

func return_one_value(value string) string {
	return value
}

func return_multiple_values() (string, string, string) {
	return "a", "b", "c"
}

func return_multiple_values_with_namedReturn() (d, e, f string) {
	d = "a"
	e = "b"
	f = "c"

	return // No explicit return value is needed, but return is needed

}
