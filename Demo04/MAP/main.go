package main

//map
import "fmt"

func main() {
	var m1 map [string]int
	fmt.Println(m1 == nil)//map[] 未初始化
	m1 = make(map[string]int,10)//估算map容量
	m1["张三"] = 18
	m1["李四"] = 20
	fmt.Println(m1) // map[张三:18 李四:20]
	fmt.Println(m1["张三"])//18
}