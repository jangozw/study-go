package main

import (
	"fmt"
)

type structTest1 struct {
	Id int
}

func main() {
	//
	var b *int
	var c int
	c = 100
	a := b
	b = &c
	fmt.Println(b, *b, a)

	//
	var d int
	var e *int
	e = &d
	d = 12
	fmt.Println(e, *e, d)

	//
	var m map[int]int
	var m2 map[int]int
	m = make(map[int]int)
	m[0] = 1
	m[1] = 5
	m2 = m
	fmt.Println(m, m2, m[1], m2[1])

	//
	var st structTest1
	var st3 *structTest1
	st2 := st
	st3 = &st
	st.Id = 2
	fmt.Println(st.Id, st2.Id, st3.Id)
}

func testPoint4() *int {
	fmt.Println("")
	a := 1
	return &a
}
