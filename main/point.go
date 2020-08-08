package main

import (
	"fmt"
)

var s5 []int

type Animaltest struct {
	Id int
}

type eat interface {
	eat() string
}

// 指针相关
func main() {
	n := 100
	a := Addr(&n) // 取n的指针地址, 作为参数传递，应为addr定义是需要指针类型
	// 返回的是一个指针地址了类型，要还原为其值，用*a
	fmt.Println(a, *a)

	s1 := make([]int, 0, 10)
	s1 = append(s1, 19)
	fmt.Println(s1)
	testS(s1)
	fmt.Println(s1, len(s1), cap(s1))
	var arr [5]int
	arr[0] = 1
	fmt.Println(arr)

	// var s2 []int;
	// s3:= make([]int,0)
	s4 := make([]int, 0)
	s4 = append(s4, 12)
	s5 := append(s4, 1)
	fmt.Println(s4, s5)

	s6 := []int{1, 2, 3}
	s7 := s6
	s7[2] = 100
	fmt.Println(s6, s7)

	m1 := make(map[string]int)
	m1["a"] = 1
	m2 := m1
	m2["c"] = 2
	fmt.Println(m1, m2)

	var f1 interface{}
	f1 = 5
	f2 := f1
	f2 = "uu"
	fmt.Println(f1, f2)

	// a := Animaltest{12}
	// b := a
	// fmt.Println(a, b)
}

func test4() {
	fmt.Println(s5)
}

func testS(s []int) {
	_ = append(s, 100)
}

// 需要一个指针类型的参数，染灰指针类型
func Addr(n *int) *int {
	b := 100 + *n // 取回n的原值
	return &b
}
