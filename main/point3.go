package main

import "fmt"

type P3 struct {
	Id   int
	Name string
}

func main() {
	var p P3
	p.Name = "hhh"

	testP3(&p)

	fmt.Println(p.Name)
}

func testP3(p *P3) {
	// p2 := *p
	// p2.Name = "8888"

	// p.Name = "www"
	cp := *p
	// fmt.Println(cp.Name)
	p = &cp
	p.Name = "999"
}
