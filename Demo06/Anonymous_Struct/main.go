package main

import "fmt"

type person struct {
	string //name
	int //age
	bool //sex
}

type student struct {
	stuAccount string
	stuInfo person
}
func main() {

	var p = person {
		"张三",
		18,
		false,
	}
	fmt.Println(p)

	var stu = student {
		stuAccount : "202000001",
		stuInfo : person{
			"李四",
			19,
			true,
		},
	}
	fmt.Println(stu)//{202000001 {李四 19 true}}


}