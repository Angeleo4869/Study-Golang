package main

//数组
import (
	"fmt"
)

func main() {
	var a [4]int//先声明后赋值
	for i := 0; i < len(a) ;i++ {
		a[i] = i
		fmt.Printf("%d ",a[i])
	}
	// a = [0 1 2 3]
	fmt.Println()
	var b = [3]int{1, 2 , 3}//直接声明且赋值
	for i := 0 ; i < 3 ;i++ {
		fmt.Printf("%d ",b[i])
	}
	// b = [1 2 3]
	fmt.Println()
	var c =[...]int{1,2,3,4,5}//根据初始值自动堆导数组大小
	for i:= 0; i<len(c); i++ {
		fmt.Printf("%d ",c[i])
	}
	// c = [1 2 3 4 5]
	fmt.Println()
	//初始化指定索引的值
	var d = [5]int{1:1 , 4:2}
	for i:= 0; i<len(d); i++ {
		fmt.Printf("%d ",d[i])
	}
	//d = [0 1 0 0 2]
	fmt.Println()


	//数组遍历 range
	var f = [4]int{1,2,3}
    for index, value := range f {
        fmt.Printf("%d %d\n",index,value)
    }
}