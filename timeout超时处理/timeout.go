package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("tome out 1")
	}

	c2 := make(chan string)

	go func() {
		c2 <- "result 2"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 1):
		fmt.Println("time out 2")
	}
}
