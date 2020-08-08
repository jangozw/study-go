package main

import (
	"context"
	"fmt"
	"time"
)

var c int

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print(i, " ")
		}()
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx Finished")
			return
		default:
			fmt.Println("doing...")
			time.Sleep(2 * time.Second)
		}
	}
	fmt.Println("All finished")
	cancel()
}

func changeCount() {
	c = 1
}
