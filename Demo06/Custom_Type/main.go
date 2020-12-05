package main

import "fmt"

type Integer int  //自定义类型
type OtherInt = int //类型别名

func main() {
	var x Integer
	var y OtherInt
	x = 10
	y = 20
	fmt.Printf("%T %v\n",x,x)//main.Integer 10
	fmt.Printf("%T %v\n",y,y)//int 20
}