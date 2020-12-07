# [Study-Golang](https://studygolang.com/pkgdoc)

## Go语言学习之路-简介

### 一、Go语言特点

1. 语法简洁
2. 开发效率高
3. 执行性能好

### 二、Go的安装

官网：[The Go Programming Language (google.cn)](https://golang.google.cn/)

下载地址：[Downloads - The Go Programming Language (google.cn)](https://golang.google.cn/dl/)

安装步骤，简单略过

### 三、设置工作路径

Go语言默认工作路径：

| 平台    | GOPATH默认值     | 举例              |
| ------- | ---------------- | ----------------- |
| Windows | %USERPROFILE%/go | C:\User\用户名\go |
| Unix    | $HOME/go         | /home/用户名/go   |

开发者可自定义工作路径，通过修改环境变量，修改工作路径

`GOPATH`=\\`自定义路径`

PATH 中添加 `%GOPATH%\bin;`

工作环境：三个文件夹 `bin`,`src`,`pkg` 

### 四、编译器

`VSCode` 安装方式略

### 五、编译方式

##### 1、使用 `go build` 编译：

1. 在项目目录下执行 ：`path >` `go build` 即可；
2. 在其他路径下执行：`path >` `go build`  `gopath` 其中  `gopath`为项目路径，且路径需从 `GOPATH/src`后开始，编译完成后，可执行文件将保存在当前目录下；
3. 重命名编译 ： `go build -o  ` ；

##### 2、使用 `go run` 编译：

该方法可在控制台直接编译执行，无需生成可执行文件；

##### 3、使用 `go install` 编译 ：

该方法等效于 先使用 方法1 编译再执行，其可执行文件保存在 `GOPATH\bin`；

### 六、跨平台编译

Go语言支持跨平台编译，（当前环境Windows ，跨平台至 Linux）

```shell
SET CGO_ENABLED=0 #禁用CGO
SET GOOS=Linux #修改为Linux平台
SET GOARCH=amd64 #设置处理器类型
```

### 七、走进Golang

1. Golang基础知识 之 变量与常量 [Demo01](/Demo01/)
2. Golang基础知识 之 程序结构 [Demo02](/Demo02/)
3. Golang基础知识 之 运算与数组 [Demo03](/Demo03/)
4. Golang基础知识 之 集合类 [Demo04](/Demo04/)
5. Golang基础知识 之 函数 [Demo05](/Demo05/)
6. Golang基础知识 之 自定义类型及结构体 [Demo06](/Demo06/)
7. Golang基础知识 之 接口与包 [Demo07](/Demo07/)
8. Golang基础知识 之 文件读写 [Demo07]