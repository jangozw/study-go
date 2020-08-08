package main

import (
	"fmt"
	"sync"
)

// 锁：lock只能lock一次，再次lock会失败，直到前面的unlock
// 互斥锁的使用，用1000个goroutine 把一个全局变量自增， 有锁跟没锁的区别，没锁肯定不准不能自增到1000，所以得有锁

var (
	testMutexCount   int
	testNoMutexCount int
)

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var mu sync.Mutex
	loopInt := 1000
	wg.Add(loopInt)
	wg2.Add(loopInt)
	for i := 0; i < loopInt; i++ {
		go countUseMutex(&wg, &mu)
		go countNoMutex(&wg2)
	}
	wg.Wait()
	wg2.Wait()
	fmt.Println("用了锁的1000个goroutine相加=", testMutexCount)
	fmt.Println("没有锁的1000个goroutine相加=", testNoMutexCount)
}

func countUseMutex(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock() // 如果多个goroutine 同时执行这里了，只有一个会加锁成功，其他的阻塞等待解锁
	testMutexCount++
	wg.Done()
	mu.Unlock()
}

func countNoMutex(wg *sync.WaitGroup) {
	testNoMutexCount++
	wg.Done()
}
