package main

import "fmt"

func main()  {
/*	// 3、new make
	var d *int  //定义一个空的指针类型的变量
	*d = 1  //对指针进行赋值,此时会报错因为没有 分配空间
	fmt.Println("d的值为：",d)*/

/*	//4、使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
	//可以直接分配地址
	a := new(int)
	//b := new(bool)
	fmt.Printf("a的类型为：%T",a)  // *int 指针类型
	fmt.Printf("a的值为：%p",a)  //一串地址 0xc00000a0b0
	fmt.Printf("对a进行指针取值操作：%v",*a)  // 0 在没有赋值的时候，会被初始化零值*/

	//5、 make(t Type, size ...IntegerType) Type
	// make也是用于内存分配的，区别于new，
	// 它只用于slice、map以及chan的内存创建，
	// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

	var m  map[string]int  //声明了一个 m变量 为map类型 key:string v:int，
	// m["测试"] = 100 // 错误的，此时只是声明，没有分配空间，不能进行赋值操作
	m = make(map[string]int,10)  //创建一个空间大小为10的map
	m["测试"] = 100
	fmt.Println(m)

	// 7、两者的区别
/*	1.二者都是用来做内存分配的。
	2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
}