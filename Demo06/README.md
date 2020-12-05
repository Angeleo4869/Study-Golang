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

结构体指针取值也可以直接使用 `指针变量.字段名` 获取字段。

5. 结构体初始化
    1. 定义时初始化   
    ```go
    type person struct {
        name string
        age int
        sex bool
    }

    func main() {
        var p1 person { //key-value 型初始化,各字段结尾必须接 ,
            name : "张三",
            age : 18,
            sex : false,
        }
        
        p2 := person {//使用值列表型初始化，但 值的顺序必须和结构体定义的字段类型的顺序一致
            "李四"
            20
            true
        }

    }
    ```

    2. [构造函数初始化](Constuctor/main.go)
    构造函数可以返回 结构体值 或 结构体指针

    一般在结构体数据较大的时候使用结构体指针，可以减小内存开销

6. 结构体内存结构
Go语言中，结构体占用一块连续的内存
```go
var t struct {
		x int
		y int
    }
    
	t.x = 10
	t.y = 20
	fmt.Printf("%T %v \n",t,t.x + t.y)//struct { x int; y int } 30 

    fmt.Printf("%p \n%p \n",&t.x,&t.y)//0xc0000140c0、0xc0000140c8
```

#### 三、[方法](Method/main.go)
Go语言中的方法 `method` 是作用于特定类型变量的函数，这种特定类型变量叫做 接收者 `Receiver` ，其类似于其他语言中的 `this` 或 `self`。
```go
func (接收者变量 接收者类型) 方法名 (参数)  (返回值) {
    函数体
}
```
接收者变量命名时，建议使用接收者类型第一个小写字母。
