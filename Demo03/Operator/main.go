package main

//运算符
import "fmt"

func main() {

	const PI = 3.14
	a := 3
	b := 4
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(b - a)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// var d float = 2 * a * PI //Go语言中，int 类型不能自动转型为 float 类型
	a++//自增1
	b--//自减1
	fmt.Println(a) 
	fmt.Println(b) 
	//在Go语言中 自增和自减 不是运算符 ，而是语句
	//比较运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	//逻辑运算符
	fmt.Println( a > 3 && b < 4)
	fmt.Println( a > 3 || b < 4)
	fmt.Println( ! (a > 3 ))
	//位运算符
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

}