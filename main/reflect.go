package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m map[reflect.Type]string
	m = make(map[reflect.Type]string)

	a := "aa"
	b := 10

	fmt.Println(reflect.TypeOf(a))

	m[reflect.TypeOf(b)] = "ad"
	m[reflect.TypeOf(a)] = "ad"
	m[reflect.TypeOf(m)] = "ad"

	fmt.Println(m)
}
