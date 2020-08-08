package main

import (
	"fmt"
	"reflect"
)

type MyError struct {
	Name string
}

func (e *MyError) Error() string {
	return "a"
}

func main() {
	// nil只能赋值给指针、channel、func、interface、map或slice类型的变量 (非基础类型) 否则会引发 panic
	var a *MyError                          // 这里不是 interface 类型  => 而是 一个值为 nil 的指针变量 a
	fmt.Printf("%+v\n", reflect.TypeOf(a))  // *main.MyError
	fmt.Printf("%#v\n", reflect.ValueOf(a)) // a => (*main.MyError)(nil)
	fmt.Printf("%p %#v\n", &a, a)           // 0xc42000c028 (*main.MyError)(nil)
	i := reflect.ValueOf(a)
	fmt.Println(i.IsNil()) // true

	if a == nil {
		fmt.Println("a == nil") //  a == nil
	} else {
		fmt.Println("a != nil")
	}

	fmt.Println("a Error():", a.Error()) // a 为什么 a 为 nil 也能调用方法？（指针类型的值为nil 也可以调用方法！但不能进行赋值操作！如下：）
	// a.Name = "1"           // panic: runtime error: invalid memory address or nil pointer dereference

	var b error = a

	// 为什么 a 为 nil 给了 b 而 b != nil ??? => error 是 interface 类型，type = *a.MyError  data = nil
	fmt.Printf("%+v\n", reflect.TypeOf(b))  // type => *main.MyError
	fmt.Printf("%+v\n", reflect.ValueOf(b)) // data => a == nil
	if b == nil {
		fmt.Println("b == nil")
	} else {
		fmt.Println("b != nil")
	}
	fmt.Println("b Error():", b.Error()) // a
}
