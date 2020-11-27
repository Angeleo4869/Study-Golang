## Demo01

#### 1、标识符与关键字：

Go语言中标识符由 数字 、字母 、及 “\_”（下划线）组成，变量名不能以 数字 及 "\_” （下划线）开头

关键字

|  **break**   |   **deafult**   |  **func**  | **interface** | **select** |
| :----------: | :-------------: | :--------: | :-----------: | :--------: |
|   **case**   |    **defer**    |   **go**   |    **map**    | **struct** |
|   **chan**   |    **else**     |  **goto**  |  **package**  | **switch** |
|  **const**   | **fallthrough** |   **if**   |   **range**   |  **type**  |
| **continue** |     **for**     | **import** |  **return**   |  **var**   |

#### 2、[变量](VAR\main.go):

声明变量：

1. `var` 声明变量

   1. ```go
      var A string  //声明变量A
      var (   //批量声明变量 B 、C 、D
          B
          C
          D
      )
      ```

   2. 声明变量并赋值

   ```go
   var str string = "Hello"
   ```

2. 类型推导

   ```GO
   var str = "Hello"
   ```

3. 简短变量

   ```go
   str := "Hello"
   ```

   简短变量声明只作用于函数内

4. 匿名变量

   匿名变量使用 ’_' 表示，不占用内存空间，不分配内存

#### 3、[常量](CONST\main.go):

1. `const` 声明常量

   ```go
   const PI = 3.1415  //声明单个常量
   const (  //批量声明常量
       A = 1
       B = 2
   )
   ```

   常量被声明后，在程序运行期间不可被修改；

   声明多个常量时，若无赋值，则当前常量值与上一个相等

2. `iota`常量计数器

`iota` 是Go语言的常量计数器，只能在常量的表达式中使用；

`iota`在`const`关键字 `const` 出现时置为 `0` ，`const`中每  <u>新增一行</u>  常量声明，将使`iota`计数一次