package main

import "fmt"

/**
区分两个概念: panic 和 报错
panic发生: 1, 手动调用panic(), 2 某些操作 如 往close 的 chan里写入数据等
报错：语法错误等，报错是报错，panic是panic
recover 只能捕获 panic
recover用法：
1， recover 用来捕获 程序中的panic错误
2, 	recover 写在defer里时候一定要把defer 放在panic语句之前才能生效捕获到
*/
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有个PANIC:", err)
		}
	}()
	b := 0
	if b == 0 {
		// 手动触发panic
		panic("被除数不能为0")
	}

	// 错误的操作 触发的panic
	ch := make(chan int)
	close(ch)
	ch <- 1 // 往一个close 的 chan 写入会panic
	fmt.Println(<-ch)

	// 触发编译错误：# command-line-arguments ， 这种跟recover 无关了
	fmt.Println(10 / 0)

	// 触发编译错误 数字超过了定义的类型范围
	var a int8
	a *= 1000
	fmt.Println(a)
}
