package main

import (
	"fmt"
	"time"
)

// 多个goroutine 执行同一个方法，对共享变量的影响

func main() {
	t1 := time.NewTicker(time.Second * time.Duration(1))
	t2 := time.NewTicker(time.Second * time.Duration(2))
	a := 1
	go func() {
		for range t1.C {
			fmt.Println("g1: ", a)
			a = 3
		}
	}()
	go func() {
		for range t2.C {
			fmt.Println("g2: ", a)
			a = 4
		}
	}()

	time.Sleep(1000 * time.Second)
}
