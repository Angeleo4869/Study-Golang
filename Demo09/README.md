### Demo09
#### 一、时间
Go语言提供时间的显示和测量函数，日历的计算方法采用的是公历。
在Go中使用时间需要导包 [`package time`](https://studygolang.com/static/pkgdoc/pkg/time.htm) 
1. 获取当前时间
```go
func Now() Time
```

2. 时间戳
`Unix()`创建一个本地时间，对应`sec`和`nsec`表示的`Unix`时间（从January 1, 1970 UTC至该时间的秒数和纳秒数）。

`nsec`的值在[0, 999999999]范围外是合法的
 

 3. 时间格式化
 Go语言中， `Time` 类型有一个自带的方法 `Format` 进行格式化，需要注意的是，Go语言中语言格式化时间模板不是常用的 `yyyy-mm-dd hh:mm:ss` 而是使用的Go的诞生时间 `2006/01/02 15:04:00` 。 

 补充： 若想使用 12 小时格式，需要指定 PM

 预定义版式
 ```go
 const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // 使用数字表示时区的RFC822
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 使用数字表示时区的RFC1123
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // 方便的时间戳
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```