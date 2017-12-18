package helper

import "fmt"

var HelperGlobalVariable string = "diffrent package - global variable"

func init() {
	fmt.Println("init helper")
}

func MethodInDifferentPackage(){    // Having the first letter capitalizes ...this method can be accessed from a diffrernet package
	fmt.Println("Method in different package")
}