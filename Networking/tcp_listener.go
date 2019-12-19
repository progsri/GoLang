package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	l, er := net.Listen("tcp", "0.0.0.0:4333")
	fmt.Println("Listening on ...", l, er)

	for true {
		data, _ := l.Accept()

		message, _ := bufio.NewReader(data).ReadString('\n')

		fmt.Println(message)

	}

	time.Sleep(100 * time.Second)
}
