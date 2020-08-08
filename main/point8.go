package main

import "fmt"

func main() {
	var a int
	var b *int
	var c **int

	b = &a
	c = &b
	a = 5

	fmt.Println(**c)
}
