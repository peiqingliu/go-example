package main

import "fmt"

func main()  {
	//1、声明一个传递int类型的通道，此时只是声明，并没有分配内存，必须在分批之后，才可用，使用make进行分配
	var ch chan int
	fmt.Println(ch) // <nil>
	// make(chan 元素类型, [缓冲大小])
	//2、分配内存
	ch1 := make(chan int)
	//3、发送数据
	//4、接收数据(从通道中接收值)(必须要从新开一个协程从通道获取，如果还是在main函数里面，会发生阻塞，且不起作用，因为main函数也是一个协程)
	//x := <- ch1
	//fmt.Println(x)
	go recv(ch1) //启用goroutine从通道接收值
	ch1 <- 10   // 把10发送到ch中
	fmt.Println("10发送成功") //使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

	//5、创建一个有缓冲的通道
	ch2 := make(chan int,2)
	ch2 <- 11
	go recv11(ch2)
	fmt.Println("11发送成功")
}

func recv(c chan int)  {
	ret := <- c
	fmt.Println("10接收到的数据：",ret)
}


func recv11(c chan int)  {
	ret := <- c
	fmt.Println("11接收到的数据：",ret)
}