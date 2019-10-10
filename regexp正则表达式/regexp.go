package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) //true

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))                   //true
	fmt.Println(r.FindString("peach punch"))              //peach
	fmt.Println(r.FindStringIndex("peach punch"))         //[0 5]
	fmt.Println(r.FindStringSubmatch("peach punch"))      //[peach ea]
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //[0 5 1 3]
	fmt.Println(r.FindAllString("peach punch pinch", -1)) //[peach punch pinch]

}
