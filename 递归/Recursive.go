package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	// fmt.Println(n * fact(n-1))
	fmt.Println(n)
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
