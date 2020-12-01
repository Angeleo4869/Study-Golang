### Demo04
#### 一、[数组](ARRAYS/main.go)：
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
3. 多维数组
    声明一个n*m的二维数组：
    ```go
    var arrays [n][m]T

    var arrays2 = [2][3]int{ //声明一个2*3 的整型数组并初始化
                {1,2,3},
                {4,5,6}
            }
    ```
注：
1、数组支持 `==`、`!=`操作符，因为内存是被初始化过的

2、`[n]*T`表示数组指针，`*[n]T`表示指针数组

#### 二、[切片](SLICE/main.go)：
切片（Slice）是一个拥有相同元素类型的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部包含`地址`、`长度`和`容量`。

切片一般用于快速地操作一块数据集合。
1. 切片的声明：
```go
var name []T
```
切片的声明与初始化方式跟数组一样，只是没有固定长度

2. 切片的长度和容量
切片拥有自己的长度和容量，可以通过 `len()`函数求长度，通过 `cap()`函数求容量
3. 基于数组定义切片
由于切片的底层就是数组，我们可以基于数组定义切片
```go
array := [...]int{1,2,3,4,5,6,7}
slice1 := array[n:m]//表示从n 到 m的切片，左闭右包
slice2 := array[:m]//表示从0 到 m的切片，左闭右包
slice3 := array[n:]//表示从n 到 end的切片，左闭右包
slice4 := array[:]//表示从0 到 end的切片，左闭右包
```
切片的容量是指底层数组的容量

底层数组，从切片的第一个元素到最后一个元素的数量

即从 `n` 开始到 `end`

切片是一个引用类型，底层数组改变，切片的值就会发生改变
4. 使用make()构造动态切片
以上切片都是基于数组来创建的，如果需要实现动态切片，需要使用内置函数`make()`，格式如下：
```go
make([]T, size, cap)
```
5. 切片的本质
切片就是一块连续内存的窗口，属于引用类型，真正的数据保存在底层数组里。

多个切片引用同一个底层数组，当其中一个值改变时，另一个也会改变

6. 切片的比较
切片之间不能直接使用`==`操作符比较，切片唯一能合法的比较的是 `nil`，一个值为`nil`的切片没有底层数组，其长度和容量都是 0 ，但是长度和容量都为 0 的切片不一定为 `nil`。

判断切片是否为空应使用`len(slice) == 0`。

7. 切片的遍历
切片遍历与数组遍历一样，可以使用索引和`for range`遍历。

8. 切片的追加
    1. append()方法为切片添加元素
    Go语言的内置函数 `append()`可以为切片添加元素
    ```go
    slice := []string{"Hello","World"}
    //slice[2] = "!" //错误的写法，会导致编译错误
    slice = append(slice,"!")//添加一个元素
    slice = append(slice,"s1","s2",...,"sn")//添加多个元素
    ```
    使用`append()`追加元素时，若原容器容量不够，则自动触发扩容机制，Go底层将开辟一块较原来容器两倍的空间，并将原容器内元素`copy`至新容器，然后在新容器追加元素，原容器将被回收。

    调用`append()`时最好用原来的变量名接受返回值。

    2. `append()`函数将元素追加到切片并返回该切片；

    3. 切片的容量以两倍进行扩容。

9. 切片扩容机制
    通过观察源码 `$GOROOT/src/runtimr/slice.go`,其中扩容相关代码如下：
    ```go
    // growslice handles slice growth during append.
    // It is passed the slice element type, the old slice, and the desired new minimum capacity,
    // and it returns a new slice with at least that capacity, with the old data
    // copied into it.
    // The new slice's length is set to the old slice's length,
    // NOT to the new requested capacity.
    // This is for codegen convenience. The old slice's length is used immediately
    // to calculate where to write new values during an append.
    // TODO: When the old backend is gone, reconsider this decision.
    // The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
    func growslice(et *_type, old slice, cap int) slice {
        //...

        newcap := old.cap
        doublecap := newcap + newcap
        if cap > doublecap {
            newcap = cap
        } else {
            if old.len < 1024 {
                newcap = doublecap
            } else {
                // Check 0 < newcap to detect overflow
                // and prevent an infinite loop.
                for 0 < newcap && newcap < cap {
                    newcap += newcap / 4
                }
                // Set newcap to the requested cap when
                // the newcap calculation overflowed.
                if newcap <= 0 {
                    newcap = cap
                }
            }
        }
    //...
    ```
    1. 首先判断，如果新申请容量`cap`大于2倍的旧容量`old.cap`，最终容量`newcap`就是新申请容量`cap`，即`(newcap = cap)`;
    2. 否则判断，如果旧切片的长度小于`1024`，则最终容量`newcap`就是旧容量`old.cap`的两倍，即`(newcap = doublecap)`；
    3. 否则判断，如果旧切片长度大于 `1024`，则最终容量`newcap`从旧容量`old.cap`开始循环增加至原来 1/4，直到最终容量`newcap`大于等于新申请容量`cap`，即
    ```go
    for 0 < newcap && newcap < cap {
		newcap += newcap / 4
    }
    ```
    4. 如果最终容量`newcap`计算值溢出，则最终容量`newcap`就是新申请容量`cap`，即`(newcap = cap)`。

需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理。

10. 切片拷贝机制
    由于切片是引用类型，所以当两个切片 a 和 b 指向同一个底层数组，其中一个修改时，另一个也会发生变化。
    Go语言内置函数 `copy()` 可以迅速地将一个切片地数据复制到另一个切片空间，其使用格式如下：
    ```go
    copy(destSlice,stcSlice []T)
    ```

11. 切片的删除
    Go语言并没有删除切片元素的专用方法，但可以通过切片的特性来删除元素
    ```go
    a = append(a[:index],a[index+1:]...)
    ```

#### 三、[指针](POINTER/main.go)：
Go语言中可以使用指针引用计算机内存地址，Go语言的取地址符号是 `&` ,根据地址取值符号是 `*`：
```go
func pointerDemo(){
    var a int = 10
    var p *int//声明一个指针变量，初始值为nil 表示没有指向任何地址
    p = &a
    fmt.Println(p,*p)
```
`&` 与 `*` 是一对互补操作符，`&`取地址，`*`根据地址取值

Go语言中不涉及到指针的运算，指针主要用于传递参数和指向内存

#### 四、 [new 与 make](NEW_MAKE/main.go)
内建函数 `new` 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针

内建函数 `make` 用来为 `slice`，`map` 或 `chan` 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，跟 `new` 类似，第一个参数也是一个类型而不是一个值，跟 `new` 不同的是，`make` 返回类型的引用而不是指针，而返回值也依赖于具体传入的类型

| 集合类型   |  描述   |
|------------|--------|
| `Slice`:  |  第二个参数 size 指定了它的长度，它的容量和长度相同,第三个参数来指定不同的容量值，但必须不能比长度值小。  |
| `Map`:     |  根据 size 大小来初始化分配内存，不过分配后的 map 长度为 0，如果 size 被忽略了，那么会在初始化分配内存时分配一个小尺寸的内存 |
| `Channel`: |  管道缓冲区依据缓冲区容量被初始化。如果容量为 0 或者忽略容量，管道是没有缓冲区的 |

#### 五、[Map](MAP/main.go)：
Go语言提供一种无序的、基于 `key - value` 数据结构的集合 `Map`，`Map`是引用类型，必须初始化才能使用
1. map定义
```go
map [keyType]valueType
```
`map`类型的变量默认初始值为 `nil`，需要使用 `make()` 函数为其分配内存
```go
make(map[keyType]valueType,[cap])
```
其中 `cap` 表示 `map` 的容量（可选），一般在初始化 `map` 时应该为其指定一个合适的初始容量

2. map遍历
`for range` 遍历 `Key - value`
```go
for k, v := range m1 {
    //~~
}
```
`for range` 遍历 `key`
```go
for k := range m1 {
    //~~~
}
```

`for range` 遍历 `value` 
```go
for _ , v := m1 {
    //~~~
}
```

3. map删除
使用 `delete()` 内置函数可以删除 `map` 中一组键值对，若`map`中不存在此`key`则不执行此操作
```go
delete(map,key)
```

#### 六、[slice 与 map 结合使用](MAP_SLICE/main.go)
`slice`与`map`在结合使用时，一定要注意初始化。