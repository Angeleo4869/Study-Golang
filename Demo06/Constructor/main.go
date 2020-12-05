
package main

import "fmt"

type person struct {
	name string
	age int
	sex bool
}
func newPerson (name string, age int, sex bool) person { //返回的是结构体值类型
	return person {
		name : name,
		age  : age,
		sex  : sex,
	}
}

func newPersonForP (name string, age int, sex bool) *person { //返回的是结构体指针
	return &person {
		name : name,
		age  : age,
		sex  : sex,
	}
}

func main() {
	fmt.Println(newPerson("张三",18,false))//{张三 18 false}
	fmt.Println(newPersonForP("李四",20,true))//&{李四 20 true}
}