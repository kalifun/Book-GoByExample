package main

import (
	"fmt"
)

func main() {
	quence := make(chan string, 2)
	quence <- "one"
	quence <- "two"
	close(quence)

	for elem := range quence {
		fmt.Println(elem)
	}
}
