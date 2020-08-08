package main

import (
	"fmt"
	"time"
)

func main() {
	// var arr [10]int{1,2,3,4,5,6,7,8,9,10}

	arr := []int{1, 2, 3, 4, 5, 6}

	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]

	fmt.Println(arr, arr1, arr2)
	resultChan := make(chan int)
	go sum(arr1, resultChan)
	go sum(arr2, resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	sum := sum1 + sum2
	fmt.Println(sum)

	orderTestOvertime()
}

func sum(arr []int, result chan int) {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	result <- sum
}

func orderTestOvertime() {
	orderChan := make(chan int32)
	timeChan := make(chan bool)
	go orderPro(orderChan)
	go overtime(timeChan)
	select {
	case <-timeChan:
		fmt.Println("Time out work")
	case <-orderChan:
		// 如何在这里得到 orderChan 的具体的值?
		// orderId := <-orderChan
		fmt.Println("Created order, id is ")
	}
}

func overtime(timeChan chan bool) {
	time.Sleep(1e9)
	timeChan <- true
}

func orderPro(orderChan chan int32) {
	var orderId int32
	orderId = 12345
	orderChan <- orderId
}
