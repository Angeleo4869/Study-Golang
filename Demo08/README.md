### 文件操作

#### 一、打开文件

1. 使用 `os.open()` 打开文件 

   `Open`打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有`O_RDONLY`模式。如果出错，错误底层类型是`*PathError`
```go
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
```

2. 使用 `os.OpenFile()` 打开文件

  `OpenFile`是一个更一般性的文件打开函数，大多数调用者都应用`Open`或`Create`代替本函数。它会使用指定的选项（如`O_RDONLY`等）、指定的模式（如`0666`等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是`*PathError`
  ```go
  func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
  ```

其中，`FileMode`为文件打开模式，其常用参数有以下几种

| 模式        | 含义                               |
| ----------- | ---------------------------------- |
| os.O_RDONLY | 只读模式打开文件                   |
| os.O_WRONLY | 只写模式打开文件                   |
| os.O_RDWR   | 读写模式打开文件                   |
| os.O_APPEND | 写操作时将数据附加到文件尾部       |
| os.O_CREATE | 如果不存在将创建一个新文件         |
| os.O_EXCL   | 和O_CREATE配合使用，文件必须不存在 |
| os.O_TRUNC  | 打开时清空文件                     |



4. Go语言中有多种读取文件的方法

   1. 使用 [`os`](https://studygolang.com/static/pkgdoc/pkg/os.htm#File) 包读取文件
   2. 使用 [`bufio`](https://studygolang.com/static/pkgdoc/pkg/bufio.htm)包读取文件
   3. 使用 [`io/ioutil`](https://studygolang.com/static/pkgdoc/pkg/io_ioutil.htm)包读取文件

5. 写入文件内容
    相应的，Go语言也提供了多种写入文件的方法
   1. 使用 [`os`](https://studygolang.com/static/pkgdoc/pkg/os.htm#File) 包写入文件
   2. 使用 [`bufio`](https://studygolang.com/static/pkgdoc/pkg/bufio.htm)包写入文件
   3. 使用 [`io/ioutil`](https://studygolang.com/static/pkgdoc/pkg/io_ioutil.htm)包写入文件