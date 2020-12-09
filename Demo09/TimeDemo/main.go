package main

import (
	"fmt"
	"time"
)

func main () {
	now := time.Now()
	fmt.Println(now)//2020-12-09 14:01:17.9251686 +0800 CST m=+0.002995501
	fmt.Println(now.Year(),now.Month(),now.Day())//2020 December 9

	fmt.Println(now.Date())//2020 December 9

	fmt.Println(now.Format(time.UnixDate))//Wed Dec  9 14:24:34 CST 2020
}