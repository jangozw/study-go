package main

import (
	"fmt"
)

/**
defer 执行的顺序：先进后出
defer 执行时候参数获取（重点）
*/
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常了", err)
		}
	}()
	var n int
	// 1,闭包： 执行的时候获取变量N的值，所以以N最后一次被赋值为准
	defer func() {
		fmt.Println("闭包内:", n)
	}()
	// 2,直接使用: defer定义时候就确定了变量，并缓纯起来等待被调用
	defer fmt.Println("直接输出 : ", n)
	// 3,函数调用:  defer定义时候就确定了变量，并缓纯起来等待被调用
	defer testDN(n)

	n = 10
	// panic 的执行顺序随机
	panic("这是一个错误")
}

func testDN(n int) {
	fmt.Println("函数内：", n)
}
