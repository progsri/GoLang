package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(array, " \n")
	fmt.Println(reflect.TypeOf(array).Kind())

	fmt.Println("--- creating a slice that references the array is came from")
	array_slice := array[2:4]
	array[3] = 222
	fmt.Print(array, " \n")
	fmt.Println(reflect.TypeOf(array).Kind())
	fmt.Print(array_slice, " \n")
	fmt.Println(reflect.TypeOf(array_slice).Kind())

	fmt.Println("---- copying the slice to another slice..this new slice does not reference the array")
	copy_array_slice := make([]int, len(array_slice))
	copy(copy_array_slice, array_slice)
	fmt.Print(copy_array_slice, " \n")
	fmt.Println(reflect.TypeOf(copy_array_slice).Kind())

	array[3] = 111
	fmt.Print(array, " \n")
	fmt.Println(reflect.TypeOf(array).Kind())

	fmt.Print(array_slice, " \n")
	fmt.Println(reflect.TypeOf(array_slice).Kind())

	fmt.Print(copy_array_slice, " \n")
	fmt.Println(reflect.TypeOf(copy_array_slice).Kind())

	array_slice = nil
	fmt.Println("--- Note # By making the first slice to nil...it will no longer reference the array and as there are no references to the array .. ",
				"the array will be garbage collected...  Memory Optimization :)")
}
