### Demo02

#### 一、[fmt占位符](FMT_Placeholder/main.go)

fmt占位符包括

| 占位符 | 作用           |
| ------ | -------------- |
| %T     | 变量属性       |
| %v     | 变量值         |
| %b     | 二进制         |
| %d     | 十进制         |
| %o     | 八进制         |
| %x     | 十六进制       |
| %s     | 字符串         |
| %#v    | 带引号的字符串 |
| %c     | 字符           |



#### 二、程序逻辑结构之判断结构 

流程控制是每种语言逻辑走向和执行次序的重要部分；

Go语言中常用的流程控制有 `if`和`for`，而`switch`和`goto`主要是为了简化代码，降低重复代码而生的结构。

1、[`if - else `](IF_ELSE/main.go)分支结构

```go
if 表达式1 {
    分支1
} else if 表达式2{
    分支2
} else {
    分支3
}
```

2、[`switch case`](SWITCH_CASE/main.go)分支结构

```go
func switchDemo(){
    finger := 3
    switch finger{
        case 条件1:{
            //~~~
        }
        case 条件2:{
            //~~~
        }
        case 条件3:{
            //~~~
        }
        case 条件4:{
            //~~~
        }
        default :{
            //~~~
        }
	}
}
```

`fallthrough`语法可以执行满足条件的`case`的下一个`case`

#### 三、程序逻辑结构之循环结构

Go语言中所有的循环类型均可以使用`for`关键字来完成。

1、[`for`](FOR/main.go)基本结构

```go
for 初始语句;条件语句;结束语句{
    循环体语句
}
```

条件表达式返回`true`时，循环体不停地循环，直到条件表达式返回`false`时自动退出循环。

for循环的初始语句和结束语句 可以忽略，例如

```go
func forDemo(){
    i := 0
    for ;i < 10 ;i++{ //忽略初始语句时，不能忽略初始语句之后的分号
        //~~~~
    }
    
    j := 0
    for j < 10 { //同时忽略初始语句和结束语句时，也可以忽略其对应的分号，类似其他编程语句的while
        //~~~
        j++
    }
}
```

无限循环：

当for循环没有写结束条件时，将会执行无限循环，此时可以通过`break`、`goto`、`return`、`panic`语句强制退出循环。

```go
for {
    循环体
}
```

for range（键值循环）

Go语言中可以使用`for range`遍历数组、切片、字符串、map及通道（channel）。通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值
2. map返回键和值
3. 通道（channel）只返回通道内的值

2、 [循环结束条件](BREAK_CONTINUE/main.go)

1. break：

   `break`关键字用于结束本次循环

   ```go
   for {
       //~~~
       if true {
           //~~~
           break;//满足条件结束循环
       }
   }
   ```

   

2. continue

   `continue`关键字用于结束当前循环，下一次循环继续执行

   ```go
   for {
       //~~~
       if true {
           //~~~
           continue  //满足条件将直接跳过本次循环
       }
       //~~~
   }
   ```

3、goto简化代码

```go
func gotoDemo() {
    for {
        //~~~
        if true {
            goto xx
        }
    }
    
    xx:
    //~~~
}
```

为了提高程序可读性，不建议使用`got`