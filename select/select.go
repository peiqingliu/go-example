package main

import (
	"fmt"
)

//多路复用
func main()  {

	output1 := make(chan string)
	output2 := make(chan string)

	//跑两个子协程
	go test1(output1)
	go test2(output2)

	//用select监控
	//select的使用类似于switch语句，
	// 它有一系列case分支和一个默认的分支。
	// 每个case会对应一个通道的通信（接收或发送）过程。
	// select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。
	select {

	case r1 := <- output1:
		fmt.Println("管道1中的数据：",r1)
	case r2 := <- output2:
		fmt.Println("管道2中的数据：",r2)
	}

}

func test1(ch chan string)  {
	ch <- "test1"
}

func test2(ch chan string)  {
	ch <- "test2"
}