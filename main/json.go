package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	OrderId int
	Username string
	Price float64
	PayTime string
}
type Request struct {
	OrderId int
}

// 结构体转json
func main()  {
	data := Response{
		OrderId: 100,
		Username: "DDDD",
		Price: 100.23,
		PayTime: "2019-09-09",
	}
	r,_ := json.Marshal(data)
	json := string(r)
	fmt.Println(json)
}