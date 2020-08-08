package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	str := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	//[]byte 转成了10进制的int数组
	byteStr := []byte(str)
	fmt.Println(byteStr)
	// 对十进制的int数组转换成字符串
	fmt.Println(string(byteStr))

	fmt.Println(strconv.ParseInt(str, 16, 32))

	s := [][]byte{[]byte("Hello"), []byte("world"), []byte("you")}
	fmt.Println(s)
	// sep := []byte("#")
	sep := []byte{}
	// join 相当于用后一个参数作为分割符，等同于php: implode(',', $arr)
	rb := bytes.Join(s, sep)
	fmt.Println(string(rb))

	// s1 :=[][]byte{[]byte("hello"), []byte("world"), []byte("fuck")}
	// s := [][]byte{[]byte("你好"),[]byte("世界")}
}
