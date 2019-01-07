package p2

import (
	"GoLang/p1/p2/p3"
	"fmt"
)

type emp p3.Employee

func (e emp) apply() int { // This would not work as the  method is defined in a different package rather than on Employee package.
	fmt.Println(e.Email)
	return 20
}

func (e emp) Apply() int {// This would not work as the Apply method is defined in a different package rather than on Employee package.
	fmt.Println(e.Email)
	return 20
}
