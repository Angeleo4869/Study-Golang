package main

//切片
import "fmt"

func main() {
	var a []int //声明一个int类型的切片

	fmt.Println(a,len(a),cap(a))//初始值为空，长度为0,容量为0

	var b = []int{1,2,3,4}//声明一个切片并初始化
	fmt.Println(b,len(b),cap(b))//初始值为[1 2 3 4]，长度为4,容量为4

	array := [...]int{1,2,3,4,5,6,7}

	c := array[1:5]//给予一个数组切割，左闭右开
	fmt.Println(c,len(c),cap(c))//初始值为[2 3 4 5] ，长度为4,容量为6



}