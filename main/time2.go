package main

import (
	"fmt"
	"time"
)

func main() {
	// timestampFormat(1568852010)

	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"

	s := [3]int{1, 2, 4}

	for _, v := range m {
		fmt.Println(v)
	}
	for _, v := range s {
		fmt.Println(v)
	}
}

func timestampFormat(t int64) {
	// timestamp := time.Now().Unix()
	// fmt.Println(timestamp)

	// 格式化为字符串,tm为Time类型

	tm := time.Unix(t, 0)

	fmt.Println(tm.Format("2006-01-02 03:04:05"))
}
