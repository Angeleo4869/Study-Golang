package main

import "fmt"
func main() { //main函数，程序的入口
	a := 2
	b := 3
	fmt.Println(a + b)
}

func sum(x int, y int)(ret int){//实现 x + y 的函数
	return x + y
}