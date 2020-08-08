package main

import (
	"fmt"
)

/**
如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
*/
func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		fmt.Println(1)
		ch <- 1
	}()
	go func() {
		fmt.Println(2)
		ch2 <- 2
	}()
	go func() {
		fmt.Println(3)
		ch3 <- 3
	}()

	// time.Sleep(100 * time.Millisecond)
	select {
	case a := <-ch:
		fmt.Println("ch", a)
	case a := <-ch2:
		fmt.Println("ch2", a)
	case a := <-ch3:
		fmt.Println("ch3", a)
	default:
		fmt.Println("default")
	}
	fmt.Println("Ok")
}
