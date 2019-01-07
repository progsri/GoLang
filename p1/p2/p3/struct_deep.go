package p3

import "fmt"

type Employee struct {
	Id           string 	//exported
	Email        string
	unexportable string    //exported
}

func (e *Employee) Skills() int {  //Exported
	fmt.Println(e.Email)
	return 10
}

func (e *Employee) skills() int {	//Unexported
	fmt.Println(e.Email)
	return 10
}

