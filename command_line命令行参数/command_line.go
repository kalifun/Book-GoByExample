package main

import (
	"fmt"
	"os"
)

func main() {
	argswithprog := os.Args
	argswithoutpro := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argswithprog)
	fmt.Println(argswithoutpro)
	fmt.Println(arg)
}
