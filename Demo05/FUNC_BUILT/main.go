package main

//内置函数 之 panic/recover
import "fmt"

func panicDemo(){
	fmt.Println("Panic Start") // 2
	defer func(){ // -2
		fmt.Println("Panic Defer")
		if r := recover(); r != nil {
			fmt.Println("Recovered From", r)
		}
		fmt.Println("Panic Exit")
	}()
	panic("Trigger Panic") //3  该语句之后程序将在执行完defer语句后结束
	// fmt.Println("Panic End")
}

func main() {
	fmt.Println("Main Start")  //1
	defer fmt.Println("Main Exit") //-1
	panicDemo() // 2
	fmt.Println("Main End")
}