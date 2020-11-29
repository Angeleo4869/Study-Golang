package main

//数组
import (
	"fmt"
)

func main() {
	var a []int = {1,2,3,4}
	for i := 0; i < len(a) ;i++ {
		fmt.Printf("%d ",a[i])
	}
}