package main

import "fmt"

func main() {
	// 在切片中追加元素
	var s []byte
	s = append(s, 20, 3)
	fmt.Println(s)

	// 合并两个切片
	var s2 []byte
	s2 = append(s2, 9, 4)
	a := append(s, s2...)
	fmt.Println(a)
}
