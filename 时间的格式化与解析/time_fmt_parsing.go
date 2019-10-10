package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339, "2018-11-06T19:57:25+08:00")
	p(t1)

	p(now.Format("3:04PM"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
}
