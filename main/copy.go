package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 5}
	s2 := make([]int, 9)
	fmt.Println(s2)
	copy(s2, s1)
	fmt.Println(s2)
}
