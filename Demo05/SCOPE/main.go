package main

//作用域
import "fmt"

var a int //全局变量

func main() {
	var b int //局部变量
	a = 5
	b = 10
	fmt.Println("a: ",a)
	fmt.Println("b: ",b)
	for i:=0; i < 10; i++ { //局部变量
		fmt.Println(f(i))
	}
}

func f(x int)(y int){//形参变量
	y = x+1
	return y
}
