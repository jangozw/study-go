package main

import (
	"fmt"
	"reflect"
	"study-go/helper"
)

func main() {
	a := helper.Stone{Name: "大理石", Height: 100}
	ra := reflect.TypeOf(a)
	fmt.Println(ra.Name(), "--", ra.Kind())

	rv := reflect.ValueOf(a)
	fmt.Println(rv.Type())
}

func TestReflectMethod() {
	person := &Person{Name: "Tom"}
	rv := reflect.ValueOf(person)
	rt := reflect.TypeOf(person)
	for i := 0; i < rt.NumMethod(); i++ {
		a := rt.Method(i)
		fmt.Println("----", a.Name, a.Func, a)
	}
	if me, ok := rt.MethodByName("GetName9"); ok {
		fmt.Println("have GetName:", me)
		value := make([]reflect.Value, 1)
		value[0] = reflect.ValueOf("hhhh")
		b := rv.MethodByName("GetName").Call(value)[0]
		fmt.Println("call method:", b)
	} else {
		fmt.Println("not found")
	}
}
