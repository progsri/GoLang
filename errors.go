package main

import "errors"
import "fmt"

func main() {

	i, e := methodRegError()
	fmt.Println(i)

	if e != nil {
		fmt.Println(e)
	}
}

func methodRegError() (int, error) {
	return 1, errors.New("returning error")
}
