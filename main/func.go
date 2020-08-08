package main

import "fmt"

func main() {
	x := 1
	y := 2
	f := func(x, y int) int {
		return x + y
	}

	b := f(x, y)

	fmt.Println(b, f)
}
