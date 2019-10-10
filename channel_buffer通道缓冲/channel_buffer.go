package main

import (
	"fmt"
)

func main() {
	message := make(chan string, 2)

	func() { message <- "kali" }()
	func() { message <- "fun" }()

	fmt.Println(<-message)
	fmt.Println(<-message)
}
