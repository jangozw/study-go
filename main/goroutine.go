package main

import (
	"fmt"
	"os"
	"sync"
)

// goroutine 执行的顺序的栈的顺序，后进先出，同defer的
// main 入口的主线程退出后，里面的goroutine也会退出，父协程退出不影响子协程继续执行

func main() {
	testGou()
}

func testGou() {
	// N次加入到队列， 有可能有输出几个，有可能没有，完全随机的；
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) // 闭包里是可以访问到i的
		}()
	}
}

func testWait() {
	var wg sync.WaitGroup
	wg.Add(2) // 要收入2个goroutine
	go func() {
		defer wg.Done() // 将wg.add 的减1
		for i := 1; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait() // 等待wg.add()个goroutine 执行完毕，否则阻塞
}

// main 入口的主线程退出后，里面的goroutine也会退出
func testGroutine001() {
	var wg sync.WaitGroup
	wg.Add(3) // 3个goroutine执行完后再推出
	for i := 0; i < 3; i++ {
		fmt.Println("counter:", i)
		go func() {
			// 是对应goroutine被执行时候取的i
			fmt.Println("goroutine:", i)
			wg.Done()
		}()
	}
	wg.Wait() // 等待goroutine执行完
}

// 父协程退出不影响子协程继续执行
func testGour() {
	go func() {
		defer fmt.Println("Dady exit")
		go func() {
			defer fmt.Printf("Son exit")
		}()
	}()
}

//
func fileAppend01() {
	f, err := os.OpenFile("/tmp/d.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o766)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < 10; i++ {
		f.WriteString(fmt.Sprintf("\n%d", i))
	}
}
