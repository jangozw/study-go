package main

import "fmt"

func main() {
	var a int
	a = 12566666666666666

	testVar(a)
	fmt.Println(a)
}

func testVar(a int) {
	a = 10
}

// func Abc() int
// func abc() int
