package main

import (
	"fmt"
)

func main() {
	a := make([]string, 3)
	fmt.Println("old:", a, "\nlen:", len(a))
	a[0] = "F"
	a[1] = "U"
	a[2] = "N"
	fmt.Println("new:", a)
	a = append(a, ",", "Best")
	fmt.Println("add:", a)

	b := make([]string, 3)
	fmt.Println("b:", b)
	copy(b, a)
	fmt.Println("b_copy:", b)

	c := a[2:4]
	fmt.Println("a-2_4:", c)

	c = a[4:]
	fmt.Println("a-4:", c)

	d := []string{"k", "a", "l", "i"}
	fmt.Println("d:", d)

	e := make([][]int, 3)
	fmt.Println("e:", e)
	for i := 0;i < 3; i++ {
		f := i+1
		e[i] = make([]int,f)
		for j := 0 ;j < f ; j++ {
			e[i][j] = i+j
		}
	}
	fmt.Println("e_final:",e)
}
