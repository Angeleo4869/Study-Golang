### Demo05
Go语言中支持函数、匿名函数和闭包
#### 一、函数定义
Go语言中定义函数使用关键字`func`：
```go
func 函数名(参数) (返回值) {
    函数体
}
```
- 函数名由字母，数字和下划线 `_` 组成，但不能以数字开头，在同一个包内，函数名不能相同
- 参数由参数变量和参数类型组成，多个参数之间用 `,`分割
```go
func f(x int, y int, str string){

}
```
- 返回值由返回值变量和变量类型组成，也可以只写返回值的类型，多个返回值必须用 `()`包裹，且用 `,` 分割
```go
func f()(x int, int, string){
    return 1, 2, "Hello"
}
```
- 函数体是实现功能的代码块
- 函数的参数与返回值是可选的，无返回值则返回值括号可以不写，无参数则参数括号必须写，返回值类型可以不命名
- 命名的返回值相当于在函数中声明一个变量，此时 `return` 后可以省略返回值
- 参数的类型简写，参数中连续相同类型的参数可以省略
```go
func f(x, y int){

}
```
- 可变长参数，其必须放在函数参数的最后
```go
func f(string , x ...int){
    //x的类型是切片
}
```

Go语言中的函数传递为值传递，及参数所传递的是其副本。

在一个命名的函数中不能声明另一个函数

#### 二、[Defer](DEFER/main.go)语句
Go语言中的 `defer` 语句会将其后面跟随的语句进行延迟处理。在 `defer` 归属的函数即将返回时，将延迟处理的语句按 `defer` 定义的逆序进行执行，即，先被 `defer` 的语句最后执行。
```go
func main() {
	fmt.Println("Start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")
}
/*
Start
End
3  
2  
1
*/
```
Go语言中函数的 `return` 不是原子操作，在底层是分两步执行
1. 返回值赋值
2. 返回返回值
函数中如果存在 `defer` ,其执行的时机在第一步和第二部之间
```go
func f1() int {//返回值 = _ = x = 5,x++,x=6,_=5
    x := 5
    defer func() {
        x++ //修改的是x不是返回值
    }()
    return x
}

func f2()(x int){ //返回值为 x = 5，x++ ，x=6
    defer func(){
        x++
    }()
    return 5
}

func f3() (y int ) { //返回值为y = x = 5 ,x++，x=6，y=5
    x := 5
    defer func() {
        x++
    }()
    return x
}

func f4 () (x int) {  //返回值为x = 5
	defer func (x int)  {
		x++ //值传递，改变的是x的副本
	}(x)
	return 5
}
```

#### 三、[作用域](SCOPE/main.go)
作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
Go 语言中变量可以在三个地方声明：
- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数
1. 局部变量
在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量.
```go
func main() {
   /* 声明局部变量 */
   var a, b, c int
   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b
   fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
}
```
2. 全局变量
在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。全局变量可以在任何函数中使用
```go
/* 声明全局变量 */
var g int

func main() {

   /* 声明局部变量 */
   var a, b int
}
```
3. 形式参数
形式参数会作为函数的局部变量来使用

4. 代码块内参数
在 `for` ，`if`， `switch` 或其他语句块内声明的变量，其作用域只能在代码块中

#### 四、[函数类型与变量](FUNC_TYPE/main.go)

函数的类型与其返回值类型一致

函数可以作为参数的类型

函数也可以作为返回值

#### 五、[匿名函数](FINC_Anonymous/main.go)
Go语言中，函数内部无法声明一个命名函数，此时可以通过匿名函数在函数内声明一个函数，
```go
func main() {
    var f = func (){
        //~~
    }
}
```
如果函数只使用一次，则可以写成立即使用函数
```go
func main() {
    func (){
        //~~
    }()
}
```
Go 语言支持匿名函数，可作为[闭包](CLOSUER/main.go)。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。

闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
```go
func sum() func(int) int {
    x := 1024
	return func(y int) int{
		return x * y
	}
}
func main() {
	f := sum()
	a := 2
    // fmt.Println("a * 1024 =",f(a))
    
}

```

#### 六、内置函数
| 内置函数 | 介绍 |
|---------|------|
|  close  | 主要用来关闭 `channel` |
|  len    | 用来求长度， `string`、`array`、`slice`、`map`、`channel` |
|  new    | 用来分配内存，主要用来分配值类型，返回的是指针 |
|  make   | 用来分配内存，主要用来分配引用类型 |
|  append | 用来追加元素到数组`array`、切片`slice`中 |
|  panic 和 recover | 用于做错误处理 |

* [Panic 和 Recover](FUNC_BUILT/main.go)
Go语言目前没有异常机制，但是可以使用 `panic/recover` 模式来处理错误。 `panic`可以在任何地方引发，但 `recover`只有在 `defer` 调用的函数中有效

```go

func panicDemo(){
	fmt.Println("Panic Start") // 2
	defer func(){ // -2
		fmt.Println("Panic Defer")
		if r := recover(); r != nil {
			fmt.Println("Recovered From", r)
		}
		fmt.Println("Panic Exit")
	}()
	panic("Trigger Panic") //3
	fmt.Println("Panic End")
}

func main() {
	fmt.Println("Main Start")  //1
	defer fmt.Println("Main Exit") //-1
	panicDemo() // 2
	fmt.Println("Main End")
}
```
panic一旦触发之后，会按照下面的顺序来做处理：

1. panic开始的地方启动终止程序操作；
2. 调用当前触发panic函数里面的defer函数；
3. 返回该函数的调用方，当作异常返回来处理，所以这一步也会调用调用方函数的defer，一直到没有调用方为止；
4. 打印panic的信息；
5. 打印堆栈跟踪信息，也就是我们看到的函数调用关系；
6. 终止程序。

recover是go提供的一个用来截获panic信息，重新获取协程控制的函数

它的使用，有两点需要注意
1. recover只能在defer函数中使用；
2. recover的使用必须与触发panic的协程是同一个协程才行

#### 七、[FMT输入输出函数](FMT_FUNC/main.go)
Go语言在实现控制台输入输出需要用到 `fmt` 包下的输入函数 `scan` 和 输出函数 `print` 其有三种输入、输出方式
```go
import "fmt"

func main() {
	var a int
	var b string
	var c float64
	//控制台输入
	fmt.Scan(&a)
	fmt.Scanf("%s",&b)//格式化输入
	fmt.Scanln(c)//输入以回车结束
	//控制台输出
	fmt.Print(a)//直接显示
	fmt.Printf("%v \n",b)//占位符显示，格式化输出
	fmt.Println(c)//单独一行显示
}
```
`scan` 与 `scanln`的区别

`scan` 函数会识别空格左右的内容，哪怕换行符号存在也不会影响 `scan` 对内容的获取

`scanln` 函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容

[其他 `fmt` 函数](https://studygolang.com/static/pkgdoc/pkg/fmt.htm)