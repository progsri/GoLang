package main

import "fmt"

func main() {

	number := 85
	switch {
	case number < 75:
		fmt.Printf("%d is less than 75 \n", number)
		fallthrough
	case number < 100:
		fmt.Printf("%d is less than 100 \n", number)
		fmt.Print(" second statement \n ")
		fallthrough  // this makes the control move to the first statement of the next case
	case number < 200:
		fmt.Printf("%d is less than 200 \n", number)
		fallthrough
	default:
		fmt.Println("DEFAULTED")
	}

}
