package main

import "fmt"
import "math"
//浮点数

func main() {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	a := 3.14
	// var b float32 = a //无法向下转换
	fmt.Printf("a:%f",a)
}