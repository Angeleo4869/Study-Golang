# Study-Golang

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

Go语言默认工作路径为（Windows环境下） C:\\User\\`username`\\go

开发者可自定义工作路径，通过修改环境变量，修改工作路径

GOPATH=\\`自定义路径`

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

Golang基础知识 之 变量与常量 [Demo01](Demo01/README.md)

