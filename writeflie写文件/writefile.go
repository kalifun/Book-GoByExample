package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\nworld\n")
	err := ioutil.WriteFile("1.txt", d1, 0644)
	check(err)

	f, err := os.Create("2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 108, 11, 5, 56}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("write %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("write %d bytes\n", n4)

	w.Flush()
}
