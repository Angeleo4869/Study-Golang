package main

//指针
import "fmt"

func main() {
	var a int = 10
    var p *int//声明一个指针变量，初始值为nil 表示没有指向任何地址
    p = &a
    
	fmt.Println(p,*p)//0xc0000140a0 10
}