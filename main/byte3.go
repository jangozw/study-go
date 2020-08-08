package main

import (
	"fmt"
)

func main() {
	var a interface{}
	var b string
	a = "AA"
	b = "sss"
	fmt.Println(a.(string), b)
	fmt.Println(a.(string), b)
}
