package main

//匿名函数
import "fmt"

func main() {
	var f1 = func(){
		fmt.Println("匿名函数f1")
	}

	f1()
	//立即执行函数，该函数之执行一次
	func (){
	fmt.Println("匿名函数")	
	}()
}