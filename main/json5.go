package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// json.Unmarshal()的第二个变量是接收参数，只能是非nil 的指针类型
// 什么样的指针类型是空指针： 只声明指针类型变量,没有赋值就是空指针
// 指向空指针的 指针类型 不是空指针。空指针是没有存放内容的指针，而指向空的指针内存已经被占用只不过存储的内容为空。
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("有个panic：", err)
		}
	}()
	testJsonNoPoint()
	testJsonEmptyPoint()
	testJsonDobuleEmptyPoint()
}

// 非指针那些不能解析，报错
func testJsonNoPoint() {
	var a int
	var b int
	b = 32
	jb, _ := json.Marshal(b)
	err := json.Unmarshal(jb, a)
	if err != nil {
		log.Println("json解析错误1:", err)
	}
	fmt.Println(a)
}

// 空指针不可以被解析, 会panic
func testJsonEmptyPoint() {
	var a *int
	c := 11
	jc, _ := json.Marshal(&c)
	err := json.Unmarshal(jc, a)
	if err != nil {
		log.Println("json解析错误2", err)
	}
	fmt.Println(*a)
}

// 指向空指针的指针可以被解析 ok
func testJsonDobuleEmptyPoint() {
	var a *int // 目前a是空指针
	var b **int
	b = &a // b 是指向空指针的指针
	c := 10
	d := &c
	jc, _ := json.Marshal(&d)
	err := json.Unmarshal(jc, b)
	if err != nil {
		log.Println("json解析错误3", err)
	}
	fmt.Println(**b)
}
