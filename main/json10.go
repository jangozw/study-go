package main

import (
	"fmt"
	"study-go/srcs/json"
)

type ApiIF interface {
	Handler()
}

type ApiIns struct {
	Name string
}

func (a *ApiIns) Handler() {
}

func main() {
	// json.Unmarshal()
	var a ApiIF
	a = &ApiIns{Name: "Dog"}
	if by, err := json.Marshal(a); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("res:", string(by))
	}
}
