package main

import (
	"database/sql/driver"
	"encoding/json"
	"time"
	//"strconv"
	"fmt"
	"strconv"
)

const YMDHIS = "2006-01-02 15:04:05" // 常规类型

type ModelTime int64

func (m ModelTime) time() time.Time {
	return time.Unix(int64(m), 0)
}

type UserInfo struct {
	Id   int
	Name string
	// CreatedAt ModelTime
	CreatedAt LocalTime
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	// 格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func main() {
	u := UserInfo{
		Id:        100,
		Name:      "2Dog",
		CreatedAt: LocalTime{time.Now()},
		// CreatedAt:1574651540,
	}
	uj, _ := json.Marshal(u)
	fmt.Println("JSON:", string(uj))

	userJson := "{\"Id\":100,\"Name\":\"2Dog\",\"CreatedAt\":1574651540}"
	var user UserInfo
	json.Unmarshal([]byte(userJson), &user)

	// fmt.Println(user.CreatedAt.time().Format(YMDHIS), user.Name)
}
