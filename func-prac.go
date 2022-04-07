package main

import (
	"fmt"
	"os"
)


func main() {
	power := power()
	fmt.Println(power)
}

func power() string {
	return os.Args[1]
}
