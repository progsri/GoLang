package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var i uint64
	// var d float64
	// var s string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner.Scan)
	fmt.Println("End")

}
