package main

import (
	"encoding/json"
	"fmt"
)

type ResJson struct {
	Code    int    `json:"code"` // 加上这个tag 在json.Marshal 时会将code 转为tag 中定义的code
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Data    string
	desc    string
}

func main() {
	// test 字段在结构体中没有所以不会设置的
	jsonStr := `{"code":0, "msg":"ok", "success":true, "test":"99000"}`
	var res ResJson

	// 将jsonStr 解析到 ResJson 结构体中去，只能解析到公有属性, 如code 对应res.Code, 不能是res.code
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)

	jsonBytes, _ := json.Marshal(res)
	jsonStr2 := string(jsonBytes)
	fmt.Println(jsonStr2)
}
