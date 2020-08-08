package main

import (
	"fmt"
	"reflect"
)

type nilSt struct {
	Id   int
	Name string
}

// 变量(非基础类型(int...))声明了没有赋值的变量 值都是nil, 这个是用 a == nil 判断的，不能fmt.Print(a) 来判断

func main() {
	var n nilSt
	var np *nilSt
	var m map[int]int
	var mp *map[int]int
	var mm []map[int]int
	var s []int
	var sp *[]int
	var c chan int
	var cp *chan int
	var a [0]int
	var ap *[2]int
	var p *int
	var i int
	var t string
	var tp *string
	var e interface{}
	var ep *interface{}
	var f func()
	var fp *func()
	mrv := reflect.ValueOf(m)
	if mrv.IsNil() {
		fmt.Println("cao ni ma", mrv.Kind())
	}

	if s == nil {
		fmt.Println("i ni ma")
	}
	e = nil
	if e == nil {
		fmt.Println("e ni ma")
	}

	fmt.Println("empty struct: ", reflect.ValueOf(n))
	fmt.Println("empty *struct: ", reflect.ValueOf(np))
	fmt.Println("empty map: ", reflect.ValueOf(m), m)
	fmt.Println("empty *map: ", reflect.ValueOf(mp))
	fmt.Println("empty []map: ", reflect.ValueOf(mm), mm)
	fmt.Println("empty slice: ", reflect.ValueOf(s))
	fmt.Println("empty *slice: ", reflect.ValueOf(sp))
	fmt.Println("empty chan int: ", reflect.ValueOf(c))
	fmt.Println("empty *chan int: ", reflect.ValueOf(cp))
	fmt.Println("empty array: ", reflect.ValueOf(a))
	fmt.Println("empty *array: ", reflect.ValueOf(ap))
	fmt.Println("empty *int: ", reflect.ValueOf(p))
	fmt.Println("empty int: ", reflect.ValueOf(i))
	fmt.Println("empty string: ", reflect.ValueOf(t))
	fmt.Println("empty *string: ", reflect.ValueOf(tp))
	fmt.Println("empty interface: ", reflect.ValueOf(e))
	fmt.Println("empty *interface: ", reflect.ValueOf(ep))
	fmt.Println("empty func: ", reflect.ValueOf(f))
	fmt.Println("empty *func: ", reflect.ValueOf(fp))

	var tea map[string]int
	var teb *map[string]int
	if teb == nil {
		fmt.Println("teb is nil hhhhh")
	}
	teb = &tea
	if teb == nil {
		fmt.Println("teb is nil")
	}
}

func testNil() (n []byte) {
	return nil
}
