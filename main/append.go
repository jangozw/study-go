package main

import "fmt"

func main() {
	a := make([]byte, 10)
	b := []byte("ab")
	c := []byte("cde")
	// a = append(a, b...)
	// a = append(a, c...)

	for i, v := range b {
		a[i] = v
	}
	for i, v := range c {
		a[i+len(b)] = v
	}

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)

	fmt.Println(3828 + 32 - 3828%32)
}
