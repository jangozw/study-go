package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()
	// 特别注意1 比如要在2的前面，chan 和 waitGroup 一起用的情况下
	// fmt.Println(<-ch)//1
	fmt.Println(<-ch) // 1
	wg.Wait()         // 2 wait 必须放到最后
}
