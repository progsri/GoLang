package loading

import (
	"GoLang/packages/flag_package"
	"fmt"
)

func main() {
	fmt.Println("loading.m1");
	Detect();
	flag_package.DisableFlag();
}


