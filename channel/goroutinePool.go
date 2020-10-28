package main

import (
	"fmt"
	"math/rand"
	"time"
)

//创建任务实例
type Job struct {
	Id int
	Random int  //随机数
}

type Result struct {
	job *Job // 任务
	sum int // 求和
}

func main()  {
	// 需要两个通道
	// 1、一个任务通道
	// 2、一个结果通道
	jobChan := make(chan *Job,128)
	resultChan := make(chan *Result,128)
	//3、创建工作池 64个协程
	createPool(64,jobChan,resultChan)
	//4、创建一个协程打印数据
	go func(resultChan chan *Result) {
		//5、从管道中取数据
		result  := <- resultChan
		fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
			result.job.Random, result.sum)
	}(resultChan)
	// 循环创建好要计算的数据
	for i :=0;i<200 ; i++ {

		job := &Job{
			Id : i,
			Random:rand.Int(), 	// 生成随机数
		}
		fmt.Println("生成的任务：",*job)
		// 把 任务放到管道里面 (go本质是通过 管道进行数据共享)
		jobChan <- job
	}

	time.Sleep(time.Microsecond)
}

/**
	num 开多少协程取跑
*/
func createPool(num int,jobChan chan *Job,resultChan chan *Result)  {
	// 开启一定数量的协程，每个协程都从管道中取读取任务
	for i :=0;i<num;i++ {
		go func(jobChan chan *Job,resultChan chan *Result) {  //匿名函数去跑
			// 1、先从管道里面读取任务
			for job := range jobChan {
				// 随机数接过来
				r_num := job.Random
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				//拼装结果
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//2、把结果扔给结果管道
				resultChan <- r
			}
		}(jobChan,resultChan)

	}
}