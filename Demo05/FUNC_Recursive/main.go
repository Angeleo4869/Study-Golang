package main

import "fmt"

func recursivDemo(x int ) int{ //递归求0 - 100 之和
	if(x >= 100 ){
		return x
	}
	return x + recursivDemo(x+1)
}
func main() {
	fmt.Println(recursivDemo(0))
}