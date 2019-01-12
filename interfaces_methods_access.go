package main

import "fmt"

type Credential struct {
	Id   string
	pass string
}

func (c Credential) valueCred() {
	fmt.Println("st value receiver ", c)
}

func (c *Credential) pointerCred() {
	fmt.Println("st pointer receiver ", c)
}

type Authorize1 interface {
	token1()
}

func (c Credential) token1() {
	fmt.Println("t1 value receiver ", c)
}

type Authorize2 interface {
	token2()
}

func (c *Credential) token2() {
	fmt.Println("t2 pointer receiver ", c)
}

func main() {
	c1 := Credential{"c1-id", "c1-pass"}
	c1.valueCred()
	c1.pointerCred()
	c1.token1()
	c1.token2()

	var i1 Authorize1
	i1 = c1
	i1.token1() // interface types can only access value methods

	var i2 Authorize2
	//i2 = c1 // this will error
	// as interface types cannot access pointer methods 

	i2 = &c1  
	i2.token2()
}
