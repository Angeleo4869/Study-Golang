package main

import "fmt"
import "encoding/json"

type person struct {
	Name string 
	Age int
}

func main() {
	p1 := person {
		Name : "张三",
		Age : 18,
	}
	p, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error:",err)
	}
	fmt.Println(string(p))//{"Name":"张三","Age":18}

	//反序列化
	str := `{"Name":"李四","Age":20}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)//传递的是指针，为了在函数内能够修改p2的值
	fmt.Println(p2)//{李四 20}
}