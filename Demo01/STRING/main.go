package main

//字符串
import "fmt"
import "strings"

func main() {
	str1 := "ABC"
	str2 := 'A'
	str3 := '汉'
	str4 := "汉字"
	str5 := "\"Hello\tWorld\n!\""  //转义字符
	str6 := `
	春 晓   
	春眠不觉晓，处处闻笛鸟。
	夜来风雨声，花落知多少`  //多行字符串
	fmt.Printf("str1: %v + %d \n" ,str1, len(str1)) //一般字符串
	fmt.Printf("str2: %v + %c \n" ,str2, str2)
	fmt.Printf("str3: %v + %c \n" ,str3, str3)
	fmt.Printf("str4: %v + %d \n" ,str4, len(str4))//汉字 及 字符串长度
	fmt.Printf("str5: %v \n" ,str5)
	fmt.Printf("%v\n",str6)

	fmt.Println(str1 + str4) // 字符串拼接
	str7 := str1 + str4// 字符串拼接
	str8 := fmt.Sprintf("%s%s",str1,str4)// 字符串拼接
	fmt.Println(str7)
	fmt.Println(str8)
	str9 := strings.Split(str7,"C") //字符串分割
	fmt.Println(str9)
	fmt.Println(strings.Contains(str7,"BC")) //判断是否含有字串
	fmt.Println(strings.HasPrefix(str7,"AB")) //判断前缀
	fmt.Println(strings.HasSuffix(str7,"汉字")) //判断后缀
	fmt.Println(strings.Index(str7,"BC") , "," , strings.LastIndex(str7,"BC")) //字串出现位置
	fmt.Println(strings.Join(strings.Split(str7,"C"),"+")) //join ,字符串拼接
}