package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM = 10000 // 求第10000位的费布拉切数

var fibs [LIM]*big.Int // 使用数组保存计算出来的费布拉切数的指针

func main() {
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		if i == LIM-1 {
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		}
	}

	fmt.Println(result)
	fmt.Printf("%T\n", result)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}
