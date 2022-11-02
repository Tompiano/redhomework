package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var ch = make(chan int64)

func add() {

	for i := 0; i < 50000; i++ {
		ch <- x
		x = x + 1
		<-ch
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
