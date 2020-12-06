package main

import "fmt"

type person struct {
	name string
	age int
	sex bool
}

//指针接收者
func newPersonForP (name string, age int, sex bool) *person { //返回的是结构体指针
	return &person {
		name : name,
		age  : age,
		sex  : sex,
	}
}

func (p *person) Birthday (){

	p.age ++
	fmt.Println(p.name,"生日啦，年龄加一岁")

}

func main() {
	var p = newPersonForP("张三",18,false)
	p.Birthday()
	fmt.Println(p)	//&{张三 19 false}
}