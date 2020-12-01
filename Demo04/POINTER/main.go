package main

//指针
import "fmt"

func main() {
	var a int = 10
	p := &a
	fmt.Println(p,*p)//0xc0000140a0 10
}