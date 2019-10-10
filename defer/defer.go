package main

import (
	"fmt"
	"os"
)

func main() {
	f := createfile("1.txt")
	defer closefile(f)
	writefile(f)
}

func createfile(p string) *os.File {
	fmt.Println("create file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writefile(f *os.File) {
	fmt.Println("write file")
	fmt.Fprintln(f, "data")
}

func closefile(f *os.File) {
	fmt.Println("close file")
	f.Close()
}
