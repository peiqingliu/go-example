package main

import (
	"fmt"
	"time"
)

func Hello()  {
	fmt.Println("Hello Goroutine!")
}

func main()  {
	//1、该示例中，执行的结果是打印完Hello Goroutine!后打印main goroutine done!。
/*	Hello()
	fmt.Println("main goroutine done!")*/
	go Hello()  // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!") //这一次的执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!。
	//2、因为此时 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。当main函数执行完毕时候，main函数内所有的程序都结束了，
	// 包括函数内的goroutine

	//3、让main函数 晚点结束
	time.Sleep(time.Second)  //这一次先打印main goroutine done!，然后紧接着打印Hello Goroutine!。


}