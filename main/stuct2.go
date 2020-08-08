package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	// 类型查询,类型断言
	// for index, data := range i {
	//	//第一个返回的是值,第二个返回的是判断内容的真假
	//	if value, ok := data.(int); ok == true {
	//		fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
	//	} else if value, ok := data.(string); ok == true {
	//		fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
	//	} else if value, ok := data.(Student); ok == true {
	//		fmt.Printf("x[%d] 类型为Student,内容为name=%s,age=%d\n", index, value.name, value.age)
	//	}
	//}

	a := Student{age: 23, name: "aaaa"}

	var b interface{}
	b = a

	if v, ok := b.(Student); ok {
		fmt.Println(v.age, v.name)
	}
}
