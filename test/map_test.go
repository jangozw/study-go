package test

import (
	"encoding/json"
	"fmt"
	"log"
	"study-go/helper"
	"testing"
)

func TestMap_Val(t *testing.T) {
	m := map[string]string{
		"a": "11",
	}
	m2 := map[string]string{
		"b": "12",
	}
	m3 := m2

	jb, _ := json.Marshal(m)
	if err := json.Unmarshal(jb, &m3); err != nil {
		log.Print(err)
	}
	fmt.Println(helper.ToJson(m3))
	fmt.Println(helper.ToJson(m2))
}
