package main

//Map 与 Slice
import "fmt"

func main() {
	//元素为Map的切片
	var s1 = make([]map[int]string,1,10) // 声明一个元素为map的切片
	s1[0] = make(map[int]string,5) //map需要初始化
	s1[0][0] = "第一个元素"
	for index, value := range s1 {
		fmt.Println("第",index,"个元素，值为",value)//第 0 个元素，值为 map[0:第一个元素]
	}

	//元素为切片类型的map
	var m1 = make(map[string][]int,10)
	m1["张三"] = []int{12,14,16}
	m1["李四"] = []int{13,15,17}
	fmt.Println(m1)//map[张三:[12 14 16] 李四:[13 15 17]]
	for k, v := range m1 {
		fmt.Println(k,": ，值为",v)//第 张三 个元素，值为 [12 14 16]  \n 第 李四 个元素，值为 [13 15 17]
	}

}