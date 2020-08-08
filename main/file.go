package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go fileAppend(&wg)

	// time.Sleep(3 * time.Second)
	fmt.Println("ok")
	wg.Wait()
}

func fileAppend(wg *sync.WaitGroup) {
	f, err := os.OpenFile("/tmp/d.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o766)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < 10; i++ {
		f.WriteString(fmt.Sprintf("\n%d", i))
	}
	wg.Done()
}

func fileWrite() {
	f, err := os.Create("a.txt")
	if err != nil {
		_ = f.Close()
		fmt.Println(err)
	}
	l, err := f.WriteString("Hello,world")
	if err != nil {
		fmt.Println(err, l)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
