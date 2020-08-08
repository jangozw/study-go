package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`href="(.*?)"`)
	html := "<a href=\"a\"></a><a href=\"b\"></a><a href=\"c\"></a><a href=\"d\"></a>"
	m := reg.FindAllStringSubmatch(html, -1)
	urls := make([]string, len(m))
	for index, value := range m {
		urls[index] = value[1]
	}
	fmt.Println(urls)
}
