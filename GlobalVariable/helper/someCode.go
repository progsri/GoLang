package helper

import "fmt"

var helperGlobalVariable string = "helper"

func init() {
	fmt.Println("init helper")
}

func methodInDifferentPackage(){
	fmt.Println("Method in different package")
}