package main

import "fmt"
//byte 和 rune 类型
func main() {
	a := 'a'
	b := '字' 
	fmt.Printf("%T + %v + %c \n",a , a , a)
	fmt.Printf("%T + %c \n",b , b)
}