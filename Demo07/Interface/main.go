package main

import "fmt"
//引出接口实例

type animal interface {
	call()//只要实现了call方法的变量都是animal类型
	eat()
}

type dog struct {}

type cat struct {}

func (d dog) call () {
	fmt.Println("狗狗叫：汪汪汪")
}

func (d dog) eat () {
	fmt.Println("狗狗吃骨头")
}

func (c cat) call () {
	fmt.Println("猫咪叫：喵喵喵")
}

func  makeAnimal (a animal) {
	a.call()
}
func main() {
	var a animal
	c := new(cat)
	d := new(dog)
	// makeAnimal(c)
	// a = c //由于猫没有实现animal的所有方法，故没有实现 animal 接口
	makeAnimal(d)
	a = d
	makeAnimal(a)

	c.call()

}