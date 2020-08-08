package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", webHanderTest)
	http.ListenAndServe("127.0.0.1:8001", nil)
}

func webHanderTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
