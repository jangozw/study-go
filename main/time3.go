package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s int
	s = 1000
	ticker := time.NewTicker(time.Second * time.Duration(1))
	go func() {
		for range ticker.C {
			s = rand.Int()
		}
	}()
	go func() {
		for {
			fmt.Println(s)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(30 * time.Second)
}
