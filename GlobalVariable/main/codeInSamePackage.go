package main

import "fmt"

var samePackageGlobalVariable = "samepackage - global variable "

func init(){
	fmt.Println("same pakage init")
}

func methodInSamePackage(){
	fmt.Println("Method in same package")
}