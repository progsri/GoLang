package main

import (
	"GoLang/packages/flag_package"
	"GoLang/packages/loading"
	"fmt"
)

func main() {
	fmt.Println("loading.m1");
	loading.Detect();
		flag_package.DisableFlag();
}