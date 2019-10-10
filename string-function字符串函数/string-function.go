package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	// Contains:  true
	p("Contains: ", s.Contains("test", "es"))

	//Count:  2
	p("Count: ", s.Count("test", "t"))

	//Hasprefix:  true
	p("Hasprefix: ", s.HasPrefix("test", "te"))

	//HasSuffix true
	p("HasSuffix", s.HasSuffix("test", "st"))

	//Index:  2
	p("Index: ", s.Index("test", "s"))

	//Join:  a-b
	p("Join: ", s.Join([]string{"a", "b"}, "-"))

	//Repeat: aaaaa
	p("Repeat:", s.Repeat("a", 5))

	//Replace: f00
	p("Replace:", s.Replace("foo", "o", "0", -1))

	//Replace: f0o
	p("Replace:", s.Replace("foo", "o", "0", 1))

	//Split: [a b c d e]
	p("Split:", s.Split("a-b-c-d-e", "-"))

	//ToLower: test
	p("ToLower:", s.ToLower("TEST"))

	//ToUpper TEST
	p("ToUpper", s.ToUpper("test"))

	//Len: 5
	p("Len:", len("hello"))

	//char: 101
	p("char:", "hello"[1])
}
