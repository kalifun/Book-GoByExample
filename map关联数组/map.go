package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a["first"] = 1
	a["second"] = 2

	fmt.Println(a)

	v1 := a["first"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(a))

	delete(a, "first")
	fmt.Println("a:", a)

	m := map[string]int{"fun": 1, "kali": 2}
	fmt.Println("m:", m)
}
