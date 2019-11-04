package main

import "fmt"

// User is very generic user
type User struct {
	name string
	id   int
	pe   Perm
}

// Perm is a nested
type Perm struct {
	allow string
	deny  string
}

func (u *User) token() *User {

	fmt.Println("value of u ", u)
	clone := u
	return clone
}

func main() {

	u := &User{name: "aa", pe: Perm{allow: "yes"}}
	fmt.Printf(" value of u %v  and %p \n", u, u)

	clone := &User{}
	*clone = *u
	fmt.Printf(" value of %v and %p ", clone, clone)
	fmt.Printf(" value of %v ", *clone)

}
