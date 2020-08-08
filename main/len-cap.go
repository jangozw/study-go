package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	// time.After(4)

	// s[0] = 200

	fmt.Println(len(s), cap(s), "--", s)

	s = append(s, 1)
	s = append(s, 2)

	fmt.Println(len(s), cap(s), "--", s)

	s = append(s, 3)
	fmt.Println(len(s), cap(s), "--", s)

	s = append(s, 4)
	fmt.Println(len(s), cap(s), "--", s)
}
