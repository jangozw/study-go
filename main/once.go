package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	a    string
	once sync.Once
)

// 这个程序还有问题
func main() {
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	// go doprint()

	time.Sleep(1e9)
}

func setup() {
	a = "Hello word"
}

func doprint() {
	once.Do(setup)
	fmt.Println(a)
	a = "No"
}
