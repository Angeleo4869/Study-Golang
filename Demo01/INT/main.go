package main

import "fmt"
import "math"

//整型
var a int = 10


func main() {
	fmt.Println(a)
	fmt.Printf("a: %b\n",a)
	fmt.Printf("a: %d\n",a)
	fmt.Printf("a: %o\n",a)
	fmt.Printf("a: %x\n",a)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
}