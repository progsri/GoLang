package main

import "fmt"


var mainGlobalVariable string = "main - global variable"

func init() {
	fmt.Println("main - init")
}

	func main(){
		fmt.Println(mainGlobalVariable)

		fmt.Println(samePackageGlobalVariable)

		methodInSamePackage()


	}