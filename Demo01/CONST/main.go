package main

import "fmt"

//常量
const PI = 3.1415926  // 声明常量

const( //批量声明常量
	OK = 200
	NOT_FOUND = 404
	N //当前常量未赋值，将与上一个常量相同
)
//iota 每声明一个常量，iota增1 
const(
	a = iota  // 0
	b = iota  // 1
	c = 100
	d 
	e = iota //4
	_ 
	f,g = iota , 100  // 6
	h = iota
)
func main()  {
	
	fmt.Println(PI)
	// PI = 3.14 //常量声明之后不能修改
	fmt.Println(OK)
	fmt.Println(NOT_FOUND)
	fmt.Println(N)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}