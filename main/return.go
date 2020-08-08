package main

import "fmt"

func main() {
	a := test()
	fmt.Println(a)
	b := test2(12)
	fmt.Println(b)
}

func test() int {
	if 1 > 2 {
		return 1
	} else {
		return 2
	}
}

func test2(x int) int {
	if x > 1 {
		return x
	} else {
		return 1
	}
}
