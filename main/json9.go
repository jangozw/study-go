package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type Per struct {
	Name string
}

func (c *Per) UnmarshalJSON(b []byte) error {
	fmt.Println("9999")
	fmt.Println(string(b))
	return nil
}

func (c Per) MarshalJSON() ([]byte, error) {
	//fmt.Println("")
	type Mper struct {
		Alias string
		Per
	}
	v := c
	fmt.Println("ok")
	m := Mper{Alias: "goo", Per: v}
	return json.Marshal(m)

}

func main() {
	/*data := &Per{Name: "aa"}
	by, _ := json.Marshal(data)
	fmt.Println(string(by))*/

	TestDec()
}

func TestDec() {
	//num := decimal.NewFromInt(88888888)
	// base := decimal.NewFromInt(11111111111)

	num := decimal.NewFromInt(int64(88888888))
	base := decimal.NewFromInt(int64(11111111111))
	// stockRight = num.DivRound(base, 7)

	n2 := num.DivRound(base, 7)
	fmt.Println(n2, n2.String())

}
