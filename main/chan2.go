package main

import "fmt"

/**
总结:
1, chan 只能在goroutine中赋值，可以赋值并不取值，不报错
2, <-chan 取值 要取goroutine里的赋的值，如果没有赋值就获取，则报异常
3，chan 被 close 后，使用<-chan 获取 其值为0
4, 给一个close的chan 赋值会异常，从一个close的chan取值不报错其值为0

*/

func main() {
	// testChan1()
	// testChan2()
	// testChan3()
	// testChan4()
	// testChan5()
	testChan6()
}

// 报错：因为chan 类型必须在goroutine中赋值
func testChan1() {
	ch := make(chan int)
	ch <- 1
}

// 报错：因为chan 类型必须在goroutine中赋值
func testChan2() {
	ch := make(chan int)
	<-ch
}

// 报错：因为chan 类型必须在goroutine中赋值，然后从外面取值
func testChan3() {
	ch := make(chan int)
	ch <- 1
	<-ch
}

// 阻塞： 因为ch 正确赋值了，却没有取出来
func testChan4() {
	ch := make(chan int)
	go func() {
		fmt.Println("test")
		ch <- 1
	}()
}

// 报错死锁：ch 只用来读取，没有在goroutine 中 赋值
func testChan5() {
	ch := make(chan int)
	go func() {
		fmt.Println("test")
	}()
	<-ch
}

// 不报错，可以给chan 赋值不读取
func testChan6() {
	ch := make(chan int)
	go func() {
		fmt.Println("test")
		ch <- 1
	}()
	a := <-ch
	fmt.Println("Hello", a)
}
