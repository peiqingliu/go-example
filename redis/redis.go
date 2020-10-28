package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("链接redis出错", err)
		return
	}
	fmt.Println("连接redis成功")
	defer c.Close()

	//设置过期时间
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
