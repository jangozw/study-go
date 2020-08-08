package main

import "fmt"

func main() {
	m := make(map[uint32]uint32)
	m[1] = 100
	var a uint32 = 1
	d, ok := m[a]
	fmt.Println(ok, d, m[a])
}
