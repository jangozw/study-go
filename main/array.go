package main

import "fmt"

func main() {
	// 数组的三种赋值形式
	var a [2]int
	a[0] = 1
	a[1] = 1

	b := [2]int{1, 2}
	d := [2]int{1, 2}
	c := [5]int{1, 2, 3, 4, 5}

	fmt.Println(c[1:3]) // 取值是 前开 后闭
	fmt.Println(a, b, c, c[2:3])
}
