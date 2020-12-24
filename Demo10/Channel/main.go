package main

//管道
import (
	"fmt"
)
func main() {
	var mychan chan int
	mychan = make(chan int,1)//定义一个管道 第二个参数为管道大小
	mychan <- 10//往管道中输入数据

	fmt.Println(mychan)
	fmt.Println(<-mychan)//从管道中读取数据
}