package main

//defer
import "fmt"

func f1() int {
    x := 5
    defer func() {
        x++
    }()
    return x
}

func f2()(x int){
    defer func(){
        x++
    }()
    return 5
}

func f3() (y int ) {
    x := 5
    defer func() {
        x++
    }()
    return x
}

func f4 () (x int) {
	defer func (x int)  {
		x++
	}(x)
	return 5
}
func main() {
	// fmt.Println("Start")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("End")

	fmt.Println("f1() = ",f1())  //5
	fmt.Println("f2() = ",f2())  //6
	fmt.Println("f3() = ",f3())  //5
	fmt.Println("f4() = ",f4())  //5
}
/*
Start
End
3  
2  
1
*/