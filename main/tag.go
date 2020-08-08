package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"username" test:"uname"`
	Age  int
	Job  string
}

// tag 用来反射时候获取属性的tag
func main() {
	p := &Person{"Fairy", 18, "teacher"}
	j, _ := json.Marshal(p)
	fmt.Printf(string(j))
	r := reflect.TypeOf(p)
	name := r.Elem().Field(0)
	fmt.Println(name.Tag.Get("json"))
	fmt.Println(name.Tag.Get("test"))
}
