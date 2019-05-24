//实现命令行参数接收
//命令行语法，下面的都可以
//eg: go run build main.go -h 127.0.0.1 -p 1000
//eg: go run build main.go --h 127.0.0.1 --p 1000
//eg: go run build main.go -h=127.0.0.1 -p=1000
//eg: go run build main.go --h=127.0.0.1 --p=1000

package main

import (
	"flag"
	"fmt"
)

//定义命令行接收的参数变量
var (
	h string
	p string
)

func main()  {
	//将h变量跟命令行的h参数绑定,设置默认值是localhost, 说明是this is host
	flag.StringVar(&h, "h", "localhost", "this is host")
	flag.StringVar(&p, "p", "0000", "this is port")
	//解析命令行参数
	flag.Parse()

	//收到参数打印出
	fmt.Println("h:", h)
	fmt.Println("p:", p)
}

