package main

import (
	"fmt"
)

/**
defer 执行顺序 ，后进先出
*/
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常了", err)
		}
	}()

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	panic("被除数不能等于0")
}
