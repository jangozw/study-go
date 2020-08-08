package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 2
	}()
	go select001(ch)
	fmt.Println("end")
}

func select001(ch chan int) {
	select {
	case ch <- 2:
		fmt.Println("ok")
	case <-time.After(1 * time.Second):
		fmt.Println("no")
	default:
		fmt.Println("default")
	}
}
