### Demo04
#### 1、[数组](ARRAYS/main.go)
数组是一种数据元素类型的集合，在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但数组大小不可改变。基本语法
```go
//定义一个长度为 n 的整型数组
var a [n]int
```
1. 数组声明
    ```go
    var 数组变量名 [元素数量]T
    ```
    数组的长度必须为常量，并且长度是数组类型的一部分。一旦定义，数组的长度不可修改。
    数组声明后，其每个元素的初始值为其对应元素类型的默认值
    数组声明并赋值
    ```go
    var array = [n]T{1,2,3,...,n}
    ```
    声明数组时可以根据初始值自动判定数组长度
    ```go
    var array = [...]T{1,2,3,...,n}
    len := len(array) //len = n
    ```
    根据索引初始化对应元素值
    ```go
    var array = [5]int{0 : 1 , 4 : 2}
    //array = { 1, 0, 0, 0, 2}
    ```
2. 数组遍历
    1. for循环遍历
    ```go
    var a = [4]int{1,2,3}
	for i := 0; i < len(a) ;i++ {
		fmt.Printf("%d ",a[i])
    }
    ```
    2. for range 遍历
    ```go
    var a = [4]int{1,2,3}
    for index, value := range a {
        fmt.Printf("%d %d\n",index,value)
    }
    ```