package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{2, 3}
	fmt.Printf("%v\n", p)    //{2 3}
	fmt.Printf("%+v\n", p)   //{x:2 y:3}
	fmt.Printf("%#v\n", p)   //main.point{x:2, y:3}
	fmt.Printf("%T\n", p)    //main.point
	fmt.Printf("%t\n", p)    //{%!t(int=2) %!t(int=3)}
	fmt.Printf("%d\n", p)    //{2 3}
	fmt.Printf("%t\n", true) //true
	fmt.Printf("%d\n", 123)  //123
	fmt.Printf("%b\n", 14)   //1110
	fmt.Printf("%c\n", 33)   //!
	fmt.Printf("%x\n", 456)  //1c8
	fmt.Printf("%f\n", 78.9) //78.900000
}
