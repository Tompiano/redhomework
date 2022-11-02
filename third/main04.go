package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		//错因：在循环体中启动协程且协程会使用循环变量的时候，所有协程
		//使用的循环变量都有可能被改写，甚至有可能在循环结束后才开始执行
		//这也是为啥输出的数有很多9
		a := i //通过显示地绑定将循环变量作为函数参数传递给协程
		go func() {
			fmt.Println(a)
			if a == 9 { //错因：使用goroutine时不给足够时间完成，main()函数直接退出。
				over <- true //所以将发送的channel放在里里面
			}
		}()

	}
	<-over
	fmt.Println("over!!!")
}
