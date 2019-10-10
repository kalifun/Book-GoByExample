package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unix := now.Unix()
	unixno := now.UnixNano()

	fmt.Println(now) //2018-11-06 19:47:13.6806177 +0800 CST m=+0.004000001
	millis := unixno / 1000000
	fmt.Println(unix)   //1541504833
	fmt.Println(millis) //1541504833680
	fmt.Println(unixno) //1541504833680617700

	fmt.Println(time.Unix(unix, 0))   //2018-11-06 19:47:13 +0800 CST
	fmt.Println(time.Unix(0, unixno)) //2018-11-06 19:47:13.6806177 +0800 CST

}
