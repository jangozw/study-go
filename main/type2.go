package main

import (
	"fmt"
	"reflect"
)

// import "fmt"

func main() {
	var a interface{}
	a = 2

	fmt.Println(reflect.TypeOf(a))

	var b []byte

	fmt.Println(len(b))
}
