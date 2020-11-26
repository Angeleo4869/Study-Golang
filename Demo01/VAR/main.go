package main

import "fmt"
//声明变量
var a int // 0
var b string // ""
var c bool // false

var ( //批量声明变量
	d int
	e string
	f bool
)
func main()  {
	
	a = 0
	b = "b + Hello"
	c = true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	// var g string  //函数内声明变量却不使用会编译错误 main.go:15:5: g declared but not used
	var h string = "h + Hello" //声明变量并同时为其赋值
	fmt.Println(h)
	var i = "i + Hello"//声明变量可以根据赋值类型推导
	fmt.Println(i)
	j := "j + Hello"//简短声明变量（只作用于函数内）
	fmt.Println(j)
	// "-" 匿名变量，

}