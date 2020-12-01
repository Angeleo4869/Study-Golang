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
	d := c[2:] //二次切片
	fmt.Println(d,len(d),cap(d))//初始值为[4 5] ，长度为2,容量为4

	//make函数创建切片
	e := make([]int ,4 , 6)
	fmt.Println(e,len(e),cap(e))//初始值为[0 0 0 0]，长度为 4，容量为 6

	// append追加元素

	e  = append(e,1,2,7,8)
	fmt.Println(e,len(e),cap(e))//初始值为[0 0 0 0 1 2 7 8]，长度为 8，容量为 12

	//copy拷贝切片
	f := make([]int ,12,12)
	copy(f, e)
	fmt.Println(f,len(f),cap(f))//初始值为[0 0 0 0 1 2 7 8 0 0 0 0]，长度为 12，容量为 12

	//删除元素
	f = append(f[4:8])
	fmt.Println(f,len(f),cap(f))//初始值为[ 1 2 7 8]，长度为 4，容量为 8


}