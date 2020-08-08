package main

import "fmt"

type (
	typeA int64
	typeB int64
)

func main() {
	// var a typeA
	var a typeA = 1
	var b typeB = 1
	if a == b {
		fmt.Println("==")
	}
}

func testType1() typeA {
	var a typeA = 122
	return a
}
