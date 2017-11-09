package main

import "fmt"

func main() {

	fmt.Println("---------- Round 1 ------------")
	variable_holds_returning_func :=  func_returns_func()
	fmt.Println(variable_holds_returning_func())
	fmt.Println(variable_holds_returning_func())

         fmt.Println("---------- Round 2 ------------")
	 variable_holds_returning_func_2 :=  func_returns_func()
	 fmt.Println(variable_holds_returning_func_2())
	 fmt.Println(variable_holds_returning_func_2())

}

func func_returns_func() func() int{
	
	value := 1
	fmt.Println("outer func value ", value)
	return  func() int {
		value = value + 1
		fmt.Println("inner func value ", value)
		return value
	}
}


