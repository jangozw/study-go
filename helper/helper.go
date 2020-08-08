package helper

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Stone struct {
	Name   string
	Height int
}

//
const YMDHIS = "2006-01-02 15:04:05" // 常规类型

//
func DateNow() string {
	return time.Now().Format(YMDHIS)
}

//
func Sha256(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	enPwd := h.Sum(nil)
	// 16进制转换为字符串
	return fmt.Sprintf("%x", enPwd)
}

func ToJson(v interface{}) string {
	j, err := json.Marshal(v)
	if err != nil {
		log.Fatal("json.Marshal err ", err)
		return ""
	}
	return string(j)
}

type Helper struct {
}

func init() {
	fmt.Println("helper init")
}

func Test() {
	fmt.Println("helper test")
}

func PrintV(v interface{}) {
	fmt.Println(v)
}

func MakeSkuChainAddr() string {
	pwd := ""
	h := sha256.New()
	h.Write([]byte(pwd))
	enPwd := h.Sum(nil)
	// 16进制转换为字符串
	return fmt.Sprintf("%x", enPwd[:20])
}

// zhangwu+#123
// 下载图片

func DownloadImg(imgUrl string, savePath string) error {
	// fileName := path.Base(imgUrl)

	// 图片正则

	// 通过http请求获取图片的流文件

	resp, err := http.Get(imgUrl)

	if err != nil || resp.StatusCode != 200 {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	out, _ := os.Create("/logs/" + ".png")

	_, err = io.Copy(out, bytes.NewReader(body))

	return err
}

// 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
