package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"d", "a", "g"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{2, 5, 1}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	ss := sort.IntsAreSorted(ints)
	fmt.Println("sort:", ss)
}
