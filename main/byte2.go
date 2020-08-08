package main

import "fmt"

func main() {
	// byte的赋值
	a := "abcdef"
	fmt.Println(len(a), len([]byte(a)))
	var b [2]byte
	b[0] = 67
	b[1] = 68
	var c []byte
	var d []byte
	for _, v := range b {
		c = append(c, v)
	}
	d = b[:]

	fmt.Println(string(c), string(d))
}
