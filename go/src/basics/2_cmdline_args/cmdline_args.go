package main

import (
	"fmt"
	"os"
)

func main() {

	// The os package provides function and values for dealing with the operating system.
	// Program argument are availabe to the os.Args variable. It is a slice.
	// os.Args[0] is the name of the program itself.
	fmt.Println("Name of program: " + os.Args[0])

	fmt.Println("The arguments:")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
