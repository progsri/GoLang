package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, er := net.Listen("tcp", "0.0.0.0:4333")
	fmt.Println("Listening on ...", l, er)

	time.Sleep(100 * time.Second)
}
