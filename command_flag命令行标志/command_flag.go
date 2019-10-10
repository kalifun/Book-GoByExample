package main

import (
	"flag"
	"fmt"
)

func main() {
	argstring := flag.String("word", "foo", "a string")
	argint := flag.Int("number", 42, "a int")
	argbool := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *argstring)
	fmt.Println("number", *argint)
	fmt.Println("fork", *argbool)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
}
