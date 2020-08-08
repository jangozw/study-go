package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"study-go/helper"
)

type cat struct {
	Id       string
	DataReal string
	Data     interface{}
}

type dog struct {
	Id       string
	Data     interface{}
	DataReal adata
}

type adata struct {
	Name string
	Age  int
}

func main() {
	var a adata
	r := reflect.ValueOf(a)
	if r.IsZero() {
		fmt.Println("is nil")
	}
	// map 不是说要 make 吗？
	// var m map[string]interface{}
	//m:= make(map[string]interface{})
	//
	m := map[string]interface{}{
		"id":   "1",
		"data": adata{Name: "Tom", Age: 12},
	}
	j, _ := json.Marshal(m)
	var j1 cat
	var j2 dog
	json.Unmarshal(j, &j1)
	fmt.Println("1---", helper.ToJson(j1))

	j1d, _ := json.Marshal(j1.Data)

	fmt.Println(reflect.TypeOf(j1d), string(j1d))

	var s string
	json.Unmarshal(j1d, &s)

	fmt.Println(s)

	json.Unmarshal(j, &j2)

	fmt.Println(j1, helper.ToJson(j1))
	// fmt.Println(j2, helper.ToJson(j2))
	// fmt.Println(m, helper.ToJson(m), j1, j2, helper.ToJson(j1), helper.ToJson(j2))
}
