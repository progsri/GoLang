package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	command1 = cobra.Command {
		Use:   "run-prepared UUID",
		Short: "Run a prepared application pod in rkt",
		Long:  `Runs a previously prepared pod by its UUID`,
	}
	command2 = &cobra.Command {
		Use:   "run-prepared UUID without &",
		Short: "Run a prepared application pod in rkt",
		Long:  `Runs a previously prepared pod by its UUID`,
	}

)


func main() {
	fmt.Println("spf13")
	fmt.Println(command1.Use)
	fmt.Println(command2.Use)
}
