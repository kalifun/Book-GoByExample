package main

import (
	"fmt"
)

func req() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	netreq := req()
	fmt.Println(netreq())
	fmt.Println(netreq())

	next := req()
	fmt.Println(next())
}
