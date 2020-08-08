package main

import (
	"fmt"
	"regexp"
)

func main() {
	line := "[asdfasd]"
	matches := regexp.MustCompile(`^\[([a-zA-Z0-9_]+)\]`).FindStringSubmatch(line)

	fmt.Println(len(matches))
}
