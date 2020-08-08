package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		}
	}
	fmt.Println(sum)

	test1(1, 23, 34)
}

func test1(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
