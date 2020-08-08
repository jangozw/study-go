package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one panic:", err)
		}
	}()
	panic("panic")
}
