package main

import (
	"fmt"
	"os"
)
func main() {
	

	/** 
	*  注：
	*  在读取文件时，若使用的是相对路径(相对该文件)，需要注意可执行文件与被读取文件的相对位置，
	*  在使用 go run 命令编译时，由于工作路径不同，可能会出现文件获取失败的情况
	*  open file failed  open ./main.go: The system cannot find the file specified.
	*  此时，使用 go build 生成可执行文件在该目录下比较合适，或者使用绝对路径(相对 $GOPATH/src)
	**/
	// fileObj, err := os.Open("./main.go")//读取文件相对 .go 文件路径
	/**open file failed  open ./main.go: The system cannot find the file specified**/
	fileObj, err := os.Open("Demo08/Test.text") //文件相对 $GOPATH/src 路径
	/***os.File**/
	if err != nil {
		fmt.Println("open file failed ",err)
		return 
	}
	defer fileObj.Close() //关闭文件 
	fmt.Printf("%T",fileObj)//*os.File
}