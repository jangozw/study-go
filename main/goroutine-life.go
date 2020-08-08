package main

import (
	"fmt"
	"time"
)

func main() {
	testLife()
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}

func testLife() {
	go func() {
		defer func() {
			fmt.Println("Dad exit")
		}()
		fmt.Println("I am Dad")
		time.Sleep(1 * time.Second)
		go func() {
			defer func() {
				fmt.Println("Son exit")
			}()
			fmt.Println("I am son")
		}()
	}()
}
