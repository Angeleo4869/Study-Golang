package main

import "fmt"

func main() {
	var a int
	var b string
	var c float64
	//控制台输入
	fmt.Scan(&a)
	fmt.Scanf("%s",&b)
	fmt.Scanln(&c)
	//控制台输出
	fmt.Print(a)//直接显示
	fmt.Printf("%v \n",b)//占位符显示，格式化输出
	fmt.Println(c)//单独一行显示
}