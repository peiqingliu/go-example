package main

import (
	"fmt"
	"runtime"
)

func main()  {

	//自定义的协程
	go func(s string) {
		for i :=0;i<4 ;i++  {
			fmt.Println(s)
		}
	}("world")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
