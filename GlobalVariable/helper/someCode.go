package helper

import "fmt"

var helperGlobalVariable string = "helper"

func init() {
	fmt.Println("init helper")
}

func MethodInDifferentPackage(){
	fmt.Println("Method in different package")
}