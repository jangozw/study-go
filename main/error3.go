package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("api error")
	errLogTest(err)
}

func errLogTest(data interface{}) {
	var m string
	// 判断data 是否实现了error 接口
	if s, ok := data.(error); ok {
		// fmt.Println("ok")
		m = s.Error()
	}
	fmt.Println("=====")
	fmt.Println(m)
}

// 基础查询
func DbBaseQuery(sellerId int64) *gorm.DB {
	query := libs.Gorm
	query = query.Where("is_delete=?", "0").Where("seller_id=?", sellerId)
	return query
}
