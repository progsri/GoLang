package main

import "fmt"
import "GoLang/GlobalVariable/helper" // this also works import "../helper"

var mainGlobalVariable string = "main - global variable"

func init() {
	fmt.Println("main - init")
}

	func main(){
		fmt.Println(mainGlobalVariable)

		fmt.Println(samePackageGlobalVariable)

		methodInSamePackage()

		tmpVariable := helper.HelperGlobalVariable

		fmt.Println(tmpVariable)
		helper.MethodInDifferentPackage()
	}

	// To execute this use go run *.go