package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
)
func readFileFromOs(fileObj *os.File) { //使用 os 包下的read() 方法读取文件
	var tmp [128]byte
	n, err := fileObj.Read(tmp[:])
	if err != nil {
		return 
	}
	fmt.Printf("读取了%d个字节\n",n)
	fmt.Println(string(tmp[:n]))
}

/** 输出格式
读取了62个字节
Hello World!
测试文本，用于做文件读写测试使用
**/

func readFileFromBufio(fileObj *os.File) {
	reader := bufio.NewReader(fileObj)
	str, err := reader.ReadString('\n') //读取字符串

	if err != nil || err == io.EOF {
		return 
	}
	fmt.Println(str)
}

/***输出
**Hello World!
**/

func readFileFromIoutil() {
	fileObj, err := ioutil.ReadFile("Demo08/Test.text") //文件相对路径
	if err != nil {
		fmt.Println("open file failed ",err)
		return 
	}
	fmt.Println(string(fileObj))
}

/***输出
**Hello World!
**测试文本，用于做文件读写测试使用
**
**/

func main() {
	fileObj, err := os.Open("Demo08/Test.text") //文件相对路径
	if err != nil {
		fmt.Println("open file failed ",err)
		return 
	}
	defer fileObj.Close() //关闭文件 

	//os.read() 方法读取文件
	readFileFromOs(fileObj)

	//bufio.readXX()方法读取文件
	// readFileFromBufio(fileObj)

	//ioutil.read()方法读取文件
	// readFileFromIoutil()
	
}
