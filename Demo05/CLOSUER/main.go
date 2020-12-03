package main

//闭包
import "fmt"

func sum() func(int) int {
    x := 1024
	return func(y int) int{
		return x * y
	}
}
func main() {
	f := sum()
	a := 2
	fmt.Println("a * 1024 = ",f(a))
}
