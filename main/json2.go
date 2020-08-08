package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string
	ServerIP   string
	Msg        string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2", "msg":"2333"}]}`

	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)
	fmt.Println(reflect.TypeOf(s))
}
