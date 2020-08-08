package main

import "fmt"

type Bird struct {
	Age   int
	Color string
}

type IFly interface {
	Fly()
	Sing()
}

func (b *Bird) Fly() {
	fmt.Println("Bird flying")
}

func (b *Bird) Sing() {
	fmt.Println("Bird singing")
}

func main() {
	// 声明fly 是IFly接口的
	var fly IFly = new(Bird)
	fly.Fly()

	// 直接初始化
	fly2 := Bird{}
	fly2.Sing()
}
