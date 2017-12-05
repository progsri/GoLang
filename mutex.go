package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex = &sync.Mutex{}
	count := 0
	for i := 0; i < 100; i++ {
		go increment(&count)
	}

	time.Sleep(time.Second * 5) // wait for the above goroutines to complete

	fmt.Println("Without locks ", count)  // this count will alwaysbe less 100 due to race conditions.

	countWithLock := 0
	for i := 0; i < 100; i++ {
		go incrementWithLock(&countWithLock, mutex)
	}

	time.Sleep(time.Second * 5)  ) // wait for the above goroutines to complete
	fmt.Println("With locks ", countWithLock)

	fmt.Println("waiting for input") 
	var input string
	fmt.Scanln(&input)

}

func increment(count *int) {
	*count = *count + 1
}

func incrementWithLock(countWithLock *int, mutex *sync.Mutex) {
	mutex.Lock()
	*countWithLock = *countWithLock + 1
	mutex.Unlock()
}
