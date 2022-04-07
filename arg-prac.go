package main

import (

	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("usage go run arg-practice.go [argument]")
		os.Exit(1)
		}

	fmt.Println("It's over",os.Args[1],"!")
}
