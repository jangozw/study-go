package main

import "fmt"

// 全局变量，函数里更改其他地方都变
var gCount int

type Animal struct {
	Name string
	Aae  int
}

/**
结论：
go 中三种引用类型 map, slice, chan， 这三种类型赋值或传参 都是对同一个对应指针的操作
相当于普通类型如int ，传递的都是&int

*/

func main() {
	fmt.Println([]byte("a"))
	if 'a' <= 97 {
		fmt.Println("A==97")
	}

	// 初始化一个切片
	slice := make([]int, 3)
	fmt.Println("原始切片", slice)
	// 在某个函数中改变其值
	changeSlice(slice)
	// 会被改变
	fmt.Println("被修改slice", slice)

	// 拷贝给另一个新变量，其实底层引用的还是一个
	slice2 := slice
	slice2[2] = 900
	// slice 也被改变了
	fmt.Println(slice)

	// 全局变量也会被改变
	changeGlobalVar()
	fmt.Println(gCount)

	mapping := make(map[string]int, 3)
	fmt.Println("原map", mapping)
	changeMap(mapping)
	fmt.Println("被修改map", mapping)

	ch := make(chan int)
	changeChan(ch)
	fmt.Println(<-ch)

	var a Animal
	fmt.Println("原struct", a)
	changeStruct(&a)
	fmt.Println("修改struct", a)
}

func changeSlice(s []int) {
	s[1] = 100
}

func changeMap(m map[string]int) {
	m["test"] = 1
}

func changeStruct(s *Animal) {
	s.Name = "Tom"
}

func changeChan(ch chan int) {
	go func() {
		ch <- 100000
	}()
}

func changeGlobalVar() {
	gCount = 1234556
}
