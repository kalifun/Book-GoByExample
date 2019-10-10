package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",") //81,
	fmt.Print(rand.Intn(100))      //87
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println("\n")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
}
