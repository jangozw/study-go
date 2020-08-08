package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		fmt.Println("ok")
	}
}
