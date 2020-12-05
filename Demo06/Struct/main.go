package main

import "fmt"

type person struct {
	name string
	age int
	sex bool
}

func setPerson(per person){
	per.age ++
	per.sex = !per.sex
}

func setPersonForP(per *person){
	per.age ++
	per.sex = !per.sex
}

func main() {
	var p person
	p.name = "张三"
	p.age = 18
	p.sex = false
	fmt.Printf("%T %v \n",p,p)//main.person {张三 18 false}

	//匿名结构体
	var t struct {
		x int
		y int
	}
	t.x = 10
	t.y = 20
	fmt.Println(t.x + t.y)

	//结构体指针
	setPerson(p)
	
	fmt.Printf("%T %v \n",p,p)//main.person {张三 18 false},由于 setPerson传递的是值拷贝，并不能改变原变量 p
	setPersonForP(&p)
	fmt.Printf("%T %v \n",p,p) //main.person {张三 19 true}通过传递 p的地址，可以对p进行修改

	var p2 = new(person)
	p2.name = "李四"
	p2.age = 20
	p2.sex = true
	fmt.Printf("%T %v \n",p2,p2)//*main.person &{李四 20 true} 指针类型的变量，
	setPersonForP(p2)
	fmt.Printf("%T %v \n",p2,p2)// *main.person &{李四 21 false} 


}