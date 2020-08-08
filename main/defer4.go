package main

import "fmt"

func main() {
	df4()
}

func df4() {
	var a int
	defer fmt.Println("--a", a)

	defer func() {
		fmt.Println("defer a:", a)
	}()

	a = 5
	fmt.Println("end a", a)
}
