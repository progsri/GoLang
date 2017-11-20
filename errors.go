package main

import "errors"
import "fmt"
import "reflect"

func main() {

	fmt.Println("------------------Regular Error--------------------------")
	i, e := methodRegError()
	fmt.Println(i)

	if e != nil {
		fmt.Println(e)
		fmt.Println(reflect.TypeOf(e)) //Tells the type
	}

	fmt.Println("------------------Custom Error--------------------------")
	custom := customError{"error name", "error desc"}
	fmt.Println(custom.Error())
	fmt.Println(reflect.TypeOf(custom))

	ci, ce := methodWithCustomError()
	fmt.Println(ci)

	if ce != nil {
		fmt.Println(ce)
		fmt.Println(reflect.TypeOf(ce))
	}
}

func methodRegError() (int, error) {
	return 1, errors.New("returning error")
}

// struct to implement Error() Method of Type error
type customError struct {
	name string
	desc string
}

func (c *customError) Error() string {
	return c.name + " " + c.desc
}

func methodWithCustomError() (int, error) {
	return 1, &customError{"custom error name", "custom error desc"}
}
