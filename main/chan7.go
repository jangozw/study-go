package main

import "fmt"

type tCh struct {
	C chan int
	M map[string]int
}

func main() {
	ch := make(chan int)
	m := make(map[string]int)
	tch := tCh{C: ch, M: m}
	go func() {
		ch <- 10
	}()
	m["Id"] = 1000
	select {
	case a := <-tch.C:
		fmt.Println("ch:", a, tch.M)
	}

	m1 := make(map[string]int)
	m2 := m1
	m2["aaa"] = 12

	mt := tCh{M: m2}
	mt.M["aaa"] = 9

	fmt.Println(m1, m2)
}
