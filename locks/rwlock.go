package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := []string{"a", "b", "c"}
	for i, a := range arr {
		fmt.Println(" %v %v", i, a)
		lock := sync.RWMutex{}

		fmt.Println(lock)
		lock.RLock()
	}

}
