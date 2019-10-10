package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7;j < 10;j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("hello")
		break
	}
}