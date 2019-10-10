package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)
	<-time1.C
	fmt.Println("Time 1")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Time 2")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Time 2 stop")
	}
}
