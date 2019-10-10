package main

import (
	"fmt"
)

func standard(output int) {
	output = 0
}

func pointer(output *int) {
	*output = 0
}

func main() {
	i := 1
	fmt.Println("i:", i)

	standard(i)
	fmt.Println("standard:", i)

	pointer(&i)
	fmt.Println("pointer:", i)

	fmt.Println("pointer_add:", &i)
}
