package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p int
	var p1 *int
	var p2 **int
	p1 = &p
	p2 = &p1
	p = 10
	fmt.Println(p1, *p1, p2, *p2, **p2)

	var r int
	json2(&r)
	fmt.Println(r)
}

func json2(r interface{}) {
	a := 33
	ja, _ := json.Marshal(a)
	b := &r
	c := &b
	d := &c
	json.Unmarshal(ja, &d)
	// fmt.Println(r)
}
