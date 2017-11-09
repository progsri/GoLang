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

//Regular method ... call starts and ends sequentially...once ended
// the variables created locally are out of scope.
// Closure ... since the reference to it always exists the current value of the 
//variable will continue to exists ( in the scope ) even after the closure method is done executing.

func func_returns_func() func() int{
	
	value := 1
	fmt.Println("outer func value ", value)
	return  func() int {
		value = value + 1     // this value gets updated..will not be the save always
		fmt.Println("inner func value ", value)
		return value
	}
}


