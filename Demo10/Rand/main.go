package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup

func f1 (i int ){
	defer wg.Done()//每次执行程序计数器减一
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}
//如何判断所有的 goroutine 已经结束？

func main () {

	rand.Seed(time.Now().UnixNano()) // 2、引入随机数种子，为 Unix时间 到现在的纳秒数
	// for i := 0; i < 10; i++ {

	// 	fmt.Println(rand.Int())
	// 	fmt.Println(rand.Intn(10))// 0-10
	// 	//1、可以发现，当程序编写完毕之后，每次执行的结果都是一样的
	// 	//   这是因为随机数的种子在编写程序之后已经确定
	// }
	// fmt.Println()
/**
在程序执行时，如何判断 子goroutine已经全部执行完毕，可通过 waitgroup 
**/
	
	for i := 0; i < 5; i++ {
		go f1(i)
		wg.Add(1)// 每次创建goroutine计数器加一
	}
	wg.Wait()//等待 waitgroup 计数器为0
}

