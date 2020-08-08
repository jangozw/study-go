package main

import "fmt"

func main() {
	arr := make([]int, 1)
	setQuote(arr)
	fmt.Println(arr)
	setQuote2(arr)
	fmt.Println(arr)
}

func setQuote(arr []int) {
	arr = append(arr, 10)
}

func setQuote2(arr []int) {
	arr[0] = 100
}
