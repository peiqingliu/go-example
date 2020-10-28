package main

import "fmt"

func main()  {

	//接口是一种类型
	//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	//
	//空接口类型的变量可以存储任意类型的变量。

	//定义一个空接口
	var x  interface{}
	s :="hello,世界"
	x = s
	fmt.Print(x)
}

