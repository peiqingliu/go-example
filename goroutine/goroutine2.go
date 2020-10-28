package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 用于记录
func record(i int)  {
	defer wg.Done()  //goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main()  {
	//4、启动多个goroutine

	for i :=0;i<10 ;i++  {
		wg.Add(1)  // 启动一个goroutine就登记+1
		go record(i)
	}
	wg.Wait()  // 等待所有登记的goroutine都结束
	
}
