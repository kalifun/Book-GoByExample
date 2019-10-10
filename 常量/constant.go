package main

import (
	"fmt"
	"math"
)

const s string = "fin"

func main() {
	fmt.Println(s)

	const n = 50000

	const e = 3e20 / n

	fmt.Println(e)

	fmt.Println(int64(e))

	fmt.Println(math.Sin(n))
}
