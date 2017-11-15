package main

import "fmt"

func main(){

	m_struct := math{12, 23}
	fmt.Println("struct w/o ptr>> ", m_struct)
	fmt.Println("struct w/o ptr add >> ", m_struct.add())    

	m_struct_ptr := &math{100, 200}
	fmt.Println("struct w ptr >> ", *m_struct_ptr)
	fmt.Println("struct w ptr add >> ", m_struct_ptr.add())  // Go automatically converts the pointer to a value(struct)
}

type math struct{
	first_value int
	second_value int
}


func (m math) add() int {
	return m.first_value + m.second_value
}

func (m *math) multiply int {
	return  *m.first_value * *m.second_value
}
