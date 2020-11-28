package main

//fmt占位符
import "fmt"

func main() {
	A := true
	B := 10
	C := 3.14
	D := "Hello"
	fmt.Printf("A: %T, value : %v \n",A,A)  //%T 查看类型  %v 查看value
	fmt.Printf("B: %T, value : %v %b  %d  %o  %x\n",B,B,B,B,B,B) //%d  %b %o %x
	fmt.Printf("C: %T, value : %v \n",C,C)
	fmt.Printf("D: %T, value : %v String : %s , %#v\n",D,D,D,D)//%s字符串 %#v 带引号的字符串
}