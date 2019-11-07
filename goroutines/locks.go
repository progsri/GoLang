package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var lock = sync.Mutex{}

func race(v1 int) {
	fmt.Printf(" Value %v \n ", v1)
	lock.Lock()
	counter = counter + 1
	lock.Unlock()
}

func main() {

	start := time.Now()
	for i := 1; i < 200000; i++ {
		go race(i)
		//time.Sleep(1 * time.Second)
	}
	end := time.Now()
	total := end.Sub(start)

	time.Sleep(30 * time.Second)
	fmt.Printf("Counter %v with time taken %v", counter, total)

}
