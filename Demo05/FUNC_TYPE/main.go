package main

//函数类型

import "fmt"

func f1() {
	fmt.Println("Hello")
}

func f2() int {
	return 10
}

func f3() (int, string){
	return 10,"Hello"
}
func f4(x int, y string)(int ){
	fmt.Println(x," ",y)
	return x + len(y)
}
func main() {
	a := f1
	fmt.Printf("f1() : %T \n",a) //f1() : func()
	b := f2
	fmt.Printf("f2() : %T \n",b)  //f2() : func() int
	c := f3
	fmt.Printf("f3() : %T",c) //f3() : func() (int, string)
	fmt.Println(f4(f3())) //15
}