package main

import (
	"fmt"
	"time"
)

func hello(i int) () {
	func () {
		i++
	}()
	fmt.Println(i)
}
//程序启动之后创建一个main goroutine 
//main函数是程序的入口
func main() {

	go fmt.Println("Groutine") //go 开启一个单独的goroutine去执行
	//
	fmt.Println("hello")//当前函数已结束，但此时 goroutine 可能还未执行完毕
	time.Sleep(time.Second)//若对程序定时，则可以看到 goroutine

	for i := 0; i < 1000; i++ {
		// go fmt.Println("当前i = ", i)
		// time.Sleep(time.Second)

		// func (i int){
		// 	go fmt.Println("当前i = ", i)
		// }(i)

		go hello(i)
	}

}