package main

import "fmt"

type person struct {
	string //name
	int //age
	bool //sex
}

func (p person) getName() (string) {
	fmt.Println("Person")
	return p.string
}

type student struct {
	stuAccount string
	person
}

func (s student) getAcconutAndName () (string,string) {
	return s.stuAccount,s.string
}

func (s student) getName () (string) { //方法覆盖
	fmt.Println("Student")
	return s.string
}
func main() {

	var p = person {
		"张三",
		18,
		false,
	}
	fmt.Println(p) //{张三 18 false}
	fmt.Println(p.getName())//Person	张三
	var stu = student {
		stuAccount : "202000001",
		person : person{
			"李四",
			19,
			true,
		},
	}
	fmt.Println(stu)//{202000001 {李四 19 true}}
	fmt.Println(stu.getName()) //Student	李四
	fmt.Println(stu.getAcconutAndName()) //202000001 李四

}