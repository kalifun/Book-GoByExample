package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("source:", a)

	a[4] = 100
	fmt.Println("new:", a)
	fmt.Println("value:", a[4])

	var b [2][3]int
	fmt.Println("what:", b)
	for i := 0;i < 2 ; i++ {
		for j :=0 ; j < 3 ; j++ {
			b[i][j] = i+j
		}
	}
	fmt.Println("for_after:",b)
}
