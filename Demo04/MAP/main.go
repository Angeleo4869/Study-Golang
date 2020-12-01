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
	m1["王五"] = 20
	m1["赵六"] = 23
	m1["李四"] = 21
	delete(m1,"张三")
	fmt.Println(m1) // map[李四:21 王五:20 赵六:23]
	fmt.Println(m1["张三"])//0
	fmt.Println(m1["张三"])
}