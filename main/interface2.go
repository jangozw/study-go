package main

import "fmt"

type Pig interface {
	Eat()
	Sleep()
}

type Pig1 struct {
	Name string
}

func (p Pig1) Eat() {
	fmt.Println("pig eat", p.Name)
}

// 这里不能 p *Pig1 否则testPig1 报错
func (p Pig1) Sleep() {
	fmt.Println("pig sleep", p.Name)
}

func testPig1(p Pig) {
	p.Sleep()
}

// 结论： 一个struct 实现了某个接口,就可以定一个函数 参数类型是该接口类型， 值传实现了该接口的struct。但是实现接口方法不能用指针接收
func main() {
	p := Pig1{Name: "Kitty"}
	testPig1(p)
}
