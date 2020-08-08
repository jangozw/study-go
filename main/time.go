package main

import "fmt"

func main() {
	fmt.Println(IsOdd(2))
	fmt.Println(IsEven(2))
}

func IsOdd(i int) bool {
	return (i & 1) == 1
}

func IsEven(i int) bool {
	return (i & 1) == 0
}
