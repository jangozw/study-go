package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	// ch2 := make(chan int, 1)
	go chanTest001(ch)
	time.Sleep(1 * time.Second)
}

func chanTest001(ch chan int) {
	fmt.Println("chan没有缓冲")
	// 赋值ch, 如果主线程不读取ch且没设置缓冲，下面的就阻塞了不执行
	ch <- 1
	fmt.Println("chan有缓冲才能输出")
}
