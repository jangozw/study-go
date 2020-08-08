package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := make(chan error)
	go func() {
		fmt.Println(1)
		fmt.Println(2)
		err <- errors.New("=========err1")
		fmt.Println(3)
		time.Sleep(1 * time.Microsecond)
		fmt.Println(4)
	}()
	go func() {
		fmt.Println(5)
		fmt.Println(6)
		err <- errors.New("=========err2")
		fmt.Println(7)
		time.Sleep(1 * time.Microsecond)
		fmt.Println(8)
	}()
	select {
	case e := <-err:
		fmt.Println(e.Error())
	case <-time.After(2 * time.Second):
		fmt.Println("nothing")
	}
}
