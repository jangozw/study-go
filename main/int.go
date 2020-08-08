package main

import (
	"fmt"
)

/**
int , uint 是根据用户操作系统决定的
int8 , uint64 等这些是固定的
*/
func main() {
	var i8 int8
	i8 = 127
	fmt.Println(i8)

	var ui8 uint8
	ui8 = 255
	fmt.Println(ui8)

	var b byte
	b = 255
	fmt.Println(b)

	var sb []byte
	sb = append(sb, 1, 2, 3)

	from := fmt.Sprintf("%x\n", sb) // why?

	fmt.Println(sb, from, sb[:2], sb[1:8]) // slice[s:e] 的取值 todo？

	// 将字符串转为数组，就是一个个字符
	str := "abcABC你好"
	var strb []byte
	strb = []byte(str)
	fmt.Println(strb)

	var a int
	a = 128
	var d uint
	d = 1280999999999
	fmt.Println(a, d)

	v1 := 32 // 自动推导的数字 是 int 型
	fmt.Println(v1)
}
