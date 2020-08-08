package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	type ta interface{}
	// var a ta
	// var a interface{}
	// var a int
	type as struct {
		Name string
	}
	// var b interface{}
	// var a int
	// a = 10
	var a as

	if err := withOutput(&a); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("ok", a)
	}

	rt := reflect.TypeOf(a)

	fmt.Println("-----", rt.Kind(), rt.Name())

	if rt.Kind() == reflect.Ptr {
		if rt.Elem().Kind() != reflect.Struct {
		}
	} else {
		if rt.Kind() != reflect.Struct {
		}
	}
}

func withOutput(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() || !rv.IsValid() {
		return errors.New("绑定值无效,必须是一个指针，不能是interface类型")
	}
	// v = 1000
	return nil
}
