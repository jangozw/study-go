package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := i
		go func() {
			if a == 3 {
				panic("其中一个goroutine挂了")
			}
			fmt.Println(a)
		}()
	}

	if err := recover(); err != nil {
		fmt.Println(err)
	}

	time.Sleep(1 * time.Second)
}
