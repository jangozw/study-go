package main

import "fmt"

// defer 执行的顺序
func main() {
	defer testdefer()
}

func testdefer() {
	// defer 执行的顺序是后进先出，同栈
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 3
		}
		fmt.Println("第一个defer") // 4
	}()
	defer func() {
		fmt.Println("第二个defer") // 2
	}()
	defer func() {
		fmt.Println("第三个defer") // 1
	}()
	// 遇到panic后，程序要退出，但还有defer要执行，于是后进先出顺序执行defer, 等执行到某个defer后有捕获异常的则输出异常
	panic("有个异常啊")
	fmt.Println("hhhhhhh")
}
