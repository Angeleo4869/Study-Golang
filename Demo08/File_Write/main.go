package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
)

func writeFileFromOs(fileObj *os.File,str string) {
	fileObj.Write([]byte(str))
	fileObj.WriteString("直接添加字符串" + str)
}

func writeFileFromBufio(fileObj *os.File,str string) {
	write := bufio.NewWriter(fileObj)
	write.WriteString(str + "\n") //将字符串写入缓存
	write.Flush() //将缓存写入文件
}

func writeFileFromIoutil(str string){
	err := ioutil.WriteFile("Demo08/Test.text",[]byte(str),0666)
	if err != nil {
		return
	}

}

func main() {
	fileObj, err := os.OpenFile("Demo08/Test.text",os.O_APPEND | os.O_WRONLY | os.O_CREATE,0644) //文件相对路径
	if err != nil {
		fmt.Println("open file failed ",err)
		return 
	}
	defer fileObj.Close() //关闭文件 
	str := "文件写入测试"
	// writeFileFromOs(fileObj,"OS包方法" + str)

	writeFileFromBufio(fileObj,"Bufio包方法" + str)

	// writeFileFromIoutil("iotil包方法 " + str)
	fmt.Println("文件写入完成")

}