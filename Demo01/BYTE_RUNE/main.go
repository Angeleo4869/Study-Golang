package main
import "strings"
import "fmt"
//byte 和 rune 类型
func main() {
	a := 'a'
	b := '字' 
	fmt.Printf("%T + %v + %c \n",a , a , a)
	fmt.Printf("%T + %c \n",b , b)

	//字符操作字符串
	hello := "Hello World !"
	var hellos  = strings.Split(hello," ")
	hellos2 := []rune(hello)  //将hello 切片 
	fmt.Printf("Hello: %v\n",hello)
	fmt.Printf("Hellos: %v\n", hellos)
	hellos[0] = "你好"//修改字符串
	// hello[0] = '你' //编译异常 cannot assign to hello[0]
	fmt.Printf("Hellos: %s \n",hellos)

	fmt.Println(hellos2) //默认以ASCLL码输出
	hellos2[1] = 'E'
	fmt.Println(string(hellos2))//转换为string
}