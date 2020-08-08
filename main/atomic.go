package main

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

func main() {
	var a atomic.Value
	var v int64 = 90
	a.Store(v)
	b := a.Load()
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
}
