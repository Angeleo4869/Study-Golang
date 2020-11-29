package main

//判断结构 if - else
import "fmt"

func main() {
	if true {
		fmt.Println("满足条件")
	} else if false{
		fmt.Println("不满足条件")
	} else {
		fmt.Println("不满足条件")
	}

	if ok := true ; ok == false  {
		fmt.Println("不满足条件" , ok)

	} else {
		fmt.Println("满足条件" , ok)
	}
}