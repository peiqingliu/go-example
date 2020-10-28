package main

import "fmt"

func main()  {
	a := 10
	b := &a  //取地址操作 &就是取一个变量的指针 也就是这个变量存储的地址
	fmt.Printf("a的值为:%d a的地址为:%p\n", a, &a)
	fmt.Printf("b的值为:%p b的类型为:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println("b的地址为：", &b)  // 0xc000006028
	c :=*b // 指针取值（根据指针去内存取值）
	fmt.Println("c的值为也就是对b(类型是指针)进行取值操作",c)  //10

	// 2、取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	//通过指针可以改变原始变量的值，
	Modify1(a)
	fmt.Println(a) // 10
	Modify2(&a)
	fmt.Println(a) // 100


}

func Modify1( x int)  {
	x = 11
}

func Modify2( x *int)  {
	//因为x的类型为 指针，指针不能直接赋值，要先取值
	*x = 100
}