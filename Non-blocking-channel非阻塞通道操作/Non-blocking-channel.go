package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	message1 := make(chan string)

	go func() {
		message <- "result"
	}()
	select {
	case msg := <-message:
		fmt.Println(msg)
	default:
		fmt.Println("No message sed")
	}

	select {
	case msg1 := <-message1:
		fmt.Println(msg1)
	default:
		fmt.Println("No message sed 1")
	}
}
