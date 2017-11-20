package main

import "fmt"

func main() {

	fmt.Println("------------ reg--------------")
	//regularMethod("regular")

	go regularMethod("go routine 1")

	// to wait for the go rountine to finish execution
	fmt.Scanln("waiting for go routine to finish execution .... this is a hack.")

	fmt.Println("------------ end --------------")
}

func regularMethod(caller string) {
	for i := 0; i < 10000; i++ {
		fmt.Printf(" %s - %d \n ", caller, i)
	}
}
