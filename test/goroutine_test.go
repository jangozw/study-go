package test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestAB(t *testing.T) {
	fmt.Println("aa", 111)
}

func TestGoroutineA(t *testing.T) {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// fmt.Println(<-gen())
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			log.Println("n=", n)
			time.Sleep(time.Second * 2)
		}
	}()
	return ch
}

func TestGoroutineTime(t *testing.T) {
	timeout := make(chan bool)
	time.AfterFunc(1*time.Second, func() {
		timeout <- true
	})
	res := make(chan int64)
	go func() {
		// time.Sleep(2 * time.Second)
		// res <- 100
		time.AfterFunc(1*time.Second, func() {
			res <- 222
		})
	}()

	var n int64
	for {
		select {
		case n = <-res:
			fmt.Println("n=", n)
		case <-timeout:
			fmt.Println("timeout!")
		}
	}
	// 这里代码执行不到的
}

func DoSomething(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Omg,timeout!")
	case <-time.After(1 * time.Second):
		fmt.Println("well done")
	}
	fmt.Println("00000000")
}

// 在context 的过期时间前做什么事情
func TestWithCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go DoSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("main ctx end:", ctx.Err())
	}
}

func TestSelectFor(t *testing.T) {
	for i := 0; i < 10; i++ {
		selectCaseByRandom()
	}
}

// select detail see：
// https://draveness.me/golang/keyword/golang-select.html
func selectCaseByRandom() {
	ch := make(chan bool)
	go func() {
		ch <- true
		fmt.Println("ok")
	}()
	var a int
	select {
	case <-ch:
		a = 1
		fmt.Println("case:", a)
		break
	case <-time.After(3 * time.Second): // 3 秒超时，假如其他case 没有再3秒内读取到，就会执行这个case, 多个case同时响应则随机
		a = 2
		fmt.Println("case:", a)
	}
	// print 1 or 2, because select pick one of cases by random
	fmt.Println("end case:", a)
}

func TestSelect(t *testing.T) {
	var a string
	for i := 0; i < 30; i++ {
		if i == 5 {
			time.Sleep(2 * time.Second)
		}
		select {
		case <-time.After(1 * time.Second):
			a = "A"
		case <-time.After(1 * time.Second):
			a = "B"
		default:
			a = "D"
		}
		fmt.Print(a)
		fmt.Print("C")

	}
}

func TestGor00(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestGor01(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestGor02(t *testing.T) {
	var i int
	for {
		if i >= 10 { // i <10 的不满足条件
			break
		}
		i++
	}
	fmt.Println(i) // 10
}
