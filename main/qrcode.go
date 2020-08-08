package main

import (
	"encoding/json"
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	n, e := QR("http://www.baidu.com")
	if e != nil {
		fmt.Println("fail")
	} else {
		j, _ := json.Marshal(n)
		fmt.Println("===", string(j))
		fmt.Println(fmt.Sprintf("%x", n))
	}
}

func QR(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, 256)
}
