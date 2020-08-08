package main

import (
	"fmt"
	"reflect"
)

type myError struct {
	Name string
}

// 实现了error接口（虽然不是故意的）
func (a *myError) Error() string {
	return "test"
}

func main() {
	var a *myError
	a = &myError{"tom"}
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) //*main.myError  test

	var b error = a                                    // 为什么 两个类型不一样能赋值？
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) //*main.myError  test
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b)) //*main.myError  test
	fmt.Println(a.Name)                                // tom
	// fmt.Println(b.Name) //报错
}
