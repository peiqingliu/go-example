package main

import "fmt"

type Person struct {
	Name string
	Age int
}

//比较两个人差值，并且返回年龄大的那个人信息，和年龄差值
func Older(p1 Person,p2 Person) (p Person,diff int) {
	if p1.Age > p2.Age {
		return p1,p1.Age - p2.Age
	}
	return p2,p2.Age - p1.Age
}

func main()  {
	p1 := Person{"小明",20}
	p2 := Person{"小李",22}
	p,diff :=Older(p1,p2)

	fmt.Printf("年龄大的人是：%v;年龄差是：%d",p,diff)
}