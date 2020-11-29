package main

//for循环
import "fmt"

func main() {
	for i:= 0; i<10; i++{ //for循环
		fmt.Println("i: ",i)
	}

	j := 0
	for ;j < 5 ;j++{
		fmt.Println("J: ",j)
	}

	for j < 10 {
		fmt.Println("J: ",j)
		j++
	}
}