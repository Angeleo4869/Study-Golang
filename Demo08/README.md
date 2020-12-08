### 文件操作
Go语言支持对文件的读写操作

1. 打开文件 
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

2. 读取文件内容


3. 写入文件内容
