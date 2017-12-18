package helper

import "fmt"

var helperGlobalVariable string = "helper"

func init() {
	fmt.Println("init helper")
}

func MethodInDifferentPackage(){    // Having the first letter capitalizes ...this method can be accessed from a diffrernet package
	fmt.Println("Method in different package")
}