package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
		ch <- 10
		ch <- 109

		/*for i := 0; i < 5 ; i++ {
			ch<-i
		}*/
	}()

	// time.Sleep(1 * time.Second) // 为什么这里不行？非要在协程里面close（ch）
	// close(ch)

	if v, ok := <-ch; ok {
		fmt.Println(v)
	}
	fmt.Println("-------")
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	fmt.Println("OK")
}
