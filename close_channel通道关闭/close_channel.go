package main

import (
	"fmt"
)

func main() {
	jbos := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jbos
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("recevied all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 0; j <= 3; j++ {
		jbos <- j
		fmt.Println("sed job ", j)
	}
	close(jbos)
	fmt.Println("sent all jobs")
	<-done
}
