package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	res := sum(2, 3)
	fmt.Printf("a+b=%d", res)
}
