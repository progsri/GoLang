package main

import "fmt"

func main(){
	fmt.Println(" struct method always take a pointer for reference to the struct..so if you call the struct method directly using struct variable then in the method you can use the variable as it is as go will pass value refered by the struct pointer ...OR... if you call the struct method using a pointer to the struct then use * to the get the value as method gets the pointer to the struct. ")
	s := math{1,2}
	fmt.Println("calling using struct variable ", s.add())

	s_ptr := &s
	fmt.Println("calling using struct pointer ", s_ptr.multiply())

}

type math struct {
	first int
	second int 
}

func (m *math) add() int {
	result := m.first + m.second
	return result
}

func (m math) multiply() int {
	result :=   *m.first * *m.second
	return result
}

