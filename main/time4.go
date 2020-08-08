package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

const YMD = "2006-01-02 15:04:05" // 常规类型

type User struct {
	Id        int64
	CreatedAt JSONTime
}

type JSONTime struct {
	time.Time
}

type TimeStr string

func (t TimeStr) JSONTime() JSONTime {
	s, _ := time.Parse(YMD, reflect.ValueOf(t).String())
	j := JSONTime{s}
	return j
}

func RedisGet(key string, v interface{}) {
	createdAtFromRedis := "2019-11-23 12:00:00"
	json.Unmarshal([]byte(createdAtFromRedis), &v)
}

func main() {
	s, _ := time.Parse(YMD, "2019-09-09 12:12:00")
	// j := JSONTime{s}

	jb, _ := json.Marshal(s)
	fmt.Println("redis set A=", string(jb))
	A := "2019-09-09T12:12:00Z"
	var out time.Time
	json.Unmarshal([]byte(A), &out)
	fmt.Println("redis get A=", out.Format(YMD))
}

func main00() {
	t1 := "2019-01-08 13:50:30" // 外部传入的时间字符串

	// 时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05"                          // 常规类型
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local) // 使用parseInLocation将字符串格式化返回本地时区时间
	log.Println("--", stamp.Unix())                                 // 输出：1546926630

	fmt.Println("=", time.Unix(stamp.Unix(), 0).Format(t1))

	timeStr := time.Now().Format("2006-01-02 15:04:05") // 当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	fmt.Println(timeStr)
}
