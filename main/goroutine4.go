package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 12
	}()

	select {
	case <-ch:
		fmt.Println("ok")
	}

	fmt.Println(<-ch)
}
