package main

import (
	"fmt"
)

func main() {
	if 7%4 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if age := 18; age >= 18 {
		fmt.Println("成年")
	} else if age < 18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("未知")
	}
}
