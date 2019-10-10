package main

import (
	"fmt"
	"sort"
)

type Bylenth []string

func (s Bylenth) Len() int {
	return len(s)
}

func (s Bylenth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Bylenth) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {
	friuts := []string{"peach", "banana", "kiwi"}
	sort.Sort(Bylenth(friuts))
	fmt.Println("sort:", friuts)
}
