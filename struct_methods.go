package main

import "fmt"

func main(){
	s := math{1,2}
	fmt.Println(s.add())
}

type math struct {
	first, second int 
}

func (m *math) add() int {
	result := m.first + m.second
	return result
}

