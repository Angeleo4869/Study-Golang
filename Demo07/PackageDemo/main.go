package main

import "fmt"
import "github.com/Study-Golang/Demo07/Util"

func main() {
	var x , y int
	x = 15
	y = 10
	fmt.Println("x + y = ",Util.Add(x,y))

	fmt.Println("x - y = ",Util.Sub(x, y))
}