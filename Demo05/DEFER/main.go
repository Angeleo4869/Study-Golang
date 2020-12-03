package main

//defer
import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")
}
/*
Start
End
3  
2  
1
*/