package main

import (
	"fmt"
)

func res() (int, int) {
	return 4, 6
}

func main() {
	a, b := res()
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)

	_, c := res()
	fmt.Printf("c=%d", c)

}
