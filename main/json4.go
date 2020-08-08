package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type req1 struct {
	Id   int
	Name string
}

// 看看json 能反解析什么类型的变量
func main() {
	var m map[string]int64
	var m2 map[string]int
	m = make(map[string]int64)
	m["a"] = 2
	m["b"] = 3
	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	err := json.Unmarshal(r, &m2)
	if err != nil {
		log.Fatal(err)
	}
	// rn,_:= json.Marshal(m2)
	fmt.Println(m["a"], m2)

	//
	var s []int
	var s2 []int
	s = make([]int, 0)
	s = append(s, 6, 7)
	r2, _ := json.Marshal(s)
	json.Unmarshal(r2, &s2)
	fmt.Println(string(r2), s, s2)

	fmt.Println("-------------------")

	testJson001()
	fmt.Println("-------------------")

	testJson002()
}

// json.Unmarshal() 第二个变量一定要是非nil的指针类型的原因
func testJson001() (res *req1) {
	// res = &req1{}
	req := req1{Id: 3, Name: "Tom"}
	jb, _ := json.Marshal(req)
	json.Unmarshal(jb, &res)
	fmt.Println("Name=", res.Name)
	return res
}

func testJson002() (res *req1) {
	r := &res
	fmt.Println(r, res)
	rv := reflect.ValueOf(r)
	if rv.Kind() == reflect.Ptr {
		fmt.Println("is Pointer type")
	}
	if rv.IsNil() {
		fmt.Println("is nil")
	}
	r2 := req1{}
	fmt.Println(r2, reflect.ValueOf(r2))
	return
}
