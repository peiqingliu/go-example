package main

import (
	"fmt"
	"sync"
)
//wg用来等待程序完成
//通过设定计数器，让每个goroutine在退出前递减，直至归零时解除阻塞。
var wg sync.WaitGroup
var lock sync.Mutex // 使用sync包的Mutex类型来实现互斥锁
var x  = 0
func main()  {
	//计数加2，表示要等待两个goroutine
	wg.Add(2)
	go Add()
	go Add()
	//等待goroutine结束
	wg.Wait()
	fmt.Println("x的值：",x)  //(资源区没加锁的情况下)多次执行值会不同
}

func Add()  {
	// 加锁
	lock.Lock()
	for i :=0;i<5000 ;i++  {
		x = x + 1
	}
	lock.Unlock()
	//释放锁
	//在函数退出时调用Done来通知main函数工作已经完成
	wg.Done()
}