package main

//new 与 make
import "fmt"

func main() {
	a := new(int)
	fmt.Printf("a: %T,%v\n",a,a)//a: *int,0xc0000140a0b
	b := make([]int,2)
	fmt.Printf("b: %T,%v\n",b,b)//[]int,[0 0]
}