package main

import (
	"fmt"
	"os"
)
func main() {
	fileObj, err := os.Open("./main.go") //文件相对路径
	if err != nil {
		fmt.Println("open file failed ",err)
		return 
	}
	defer fileObj.Close() //关闭文件 
}