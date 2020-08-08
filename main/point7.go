package main

import "fmt"

func main() {
	var a int
	var b *int
	var c **int
	var d ***int
	b = &a
	c = &b
	d = &c
	// a = 13
	fmt.Println(*b, **c, ***d) // 13
}
