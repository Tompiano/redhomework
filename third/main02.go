package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan struct{})
	wg1 sync.WaitGroup
)

func main() {
	wg1.Add(2) //设置两个goroutine

	go Print1()
	Print2()
	wg1.Wait() //让主协程等子协程进行完后再退出

}

func Print1() {
	defer wg1.Done()
	//打印奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println("Print1:", i)
		ch1 <- struct{}{} //Print1输出后关闭Print1让Print2输出
		<-ch1             //开启Print1，防止通道死锁

	}

}

func Print2() {
	defer wg1.Done()

	for i := 2; i <= 100; i += 2 {
		<-ch1 //关闭Print2让Print1输出
		fmt.Println("Print2:", i)
		ch1 <- struct{}{} //开启Print2让其输出

	}

}
