package main

//循环结束
import "fmt"

func main() {
	for i := 0 ;i < 10 ;i++ {
		if(i == 5 ){//i == 5 结束循环，>5 的数将不会执行
			break
		}
		fmt.Println(i)
	}
	fmt.Println("END")
	for i := 0 ;i < 10 ;i++ {
		if(i == 5 ){//i == 5 跳过当前循环，>5 的数将依然执行
			continue
		}
		fmt.Println(i)
	}
} 