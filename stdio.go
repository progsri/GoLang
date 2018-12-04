package main

import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println("standard input ... file descriptor 0")

	scan := bufio.NewScanner(os.Stdin)

	// Takes contnous inputs
	count := 0
	for scan.Scan() {

		fmt.Println(scan.Text())
		count = count + 1

		if count == 3 {
			break
		}
	}

	fmt.Println("Done")

}
