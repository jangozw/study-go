package helper

import (
	"encoding/json"
	"fmt"
)

// Hello is test
func Hello() {
	fmt.Println("asdf", "asdf")
}

// World is just for test
func World() []byte {
	a, _ := json.Marshal([]byte("world"))
	return a
}
