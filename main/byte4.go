package main

type orderInfo struct {
	Id       string
	Buyer    int64
	CreateAt int64
	Msg      string
}

func main() {
	order := orderInfo{
		Id:       "202099999999999999",
		Buyer:    123456,
		CreateAt: 0o000,
	}
}
