### Demo06
#### 一、[自定义类型和别名](Custom_Type/main.go)
Go语言没有“类”的概念，也不支持“类”的己成等面向对象的概念。
1. 类型别名和自定义类型
Go语言中可以使用 `type` 关键字来定义自定义类型
```go
type Integer int
```

#### 二、[结构体](Struct/main.go)
Go语言中的基础数据类型可以表示一些事物的基本属性，但需要表达一个事物的部分或全部属性时，单一的基本数据类型无法满足需求，Go语言提供了一种自定义的数据类型，可以封装多个基本数据类型，称之为结构体 `struct` ，结构体变量是值类型。

1. 结构体定义
Go语言中使用 `type` 和 `struct` 关键字定义结构体
```go
type 类型名 struct {
    字段名1 字段类型
    字段名2 字段类型
    ...
}
```

2. 结构体字段
结构体变量可以通过 `.` 来访问其各个字段
```go
package main

import "fmt"

type person struct {
	name string
	age int
	sex bool
}
func main() {
	var p person
	p.name = "张三"
	p.age = 18
	p.sex = false
	fmt.Printf("%T %v \n",p,p)//main.person {张三 18 false}
}
```

3. 匿名结构体
```go
//匿名结构体
	var t struct {
		x int
		y int
	}
```
匿名结构体多用于临时场景

4. 结构体指针
在Go语言中，结构体变量为值类型，所以在进行参数传递时，所传递的是拷贝，而不是该变量本身，所以，在使用函数对某结构体变量进行操作时，需要传递指针。
```go
type T struct {
		x int
		y int
    }
    
func setT(t T){
    t.x ++
    t.y--
}

func setT(&t T){
    *t.x ++
    *t.y--
}

func main() {
    var t T
    t.x = 10
    t.y = 20
    fmt.Println(t.x,t.y)
}
