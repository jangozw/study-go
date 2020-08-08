package main

import (
	"fmt"
	"sync"
)

// 1,同时执行多个goroutine 其中一个挂了，其他的也会挂了
func main() {
	// ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		//a:=i
		go func(a int) {
			fmt.Println(a) // 1, i 这个循环变量在闭包中去的是最后一次的i，如果想要当时的i，则可以穿参数
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("ok")
}

func testChan12(ch chan int, wg *sync.WaitGroup) {
	ch <- 12
	wg.Done()
}
