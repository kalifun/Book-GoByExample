package main

import (
	"os"
)

func main() {
	panic("have a problem")

	_, err := os.Create("home/ubuntu")
	if err != nil {
		panic(err)
	}
}
