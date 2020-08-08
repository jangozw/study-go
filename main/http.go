package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var counter int

func main() {
	defer func() {
		log.Println("main exit")
	}()
	fmt.Println("Http server started")
	http.HandleFunc("/", reqHandel)

	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func reqHandel(w http.ResponseWriter, r *http.Request) {
	counter++
	log.Println("调用次数:", counter)
	go func() {
		time.Sleep(1 * time.Second)
		log.Println("一个耗时的groutine执行了")
	}()

	go func() {
		// panic("发生异常")
	}()

	args := r.URL.Query()
	fmt.Println(args.Get("a"))
	fmt.Fprintln(w, args)
	w.Write([]byte("aaaaaaaaaaa"))
}
