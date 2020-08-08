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
		fmt.Println("00999")
		ch <- 1
		wg.Done()
	}()
	//<-ch 必须取值，否则报错
	wg.Wait()
	// time.Sleep(1* time.Second) //这种情况不报错
}
