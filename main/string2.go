package main

import (
	"bytes"
	"fmt"
)

// 字符串拼接
func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	b.Write([]byte(","))
	b.Write([]byte("world"))
	fmt.Println(b.String())

	// b2 := bytes.NewBuffer([]byte{'a', 'b', 'c'})
	// buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
}
