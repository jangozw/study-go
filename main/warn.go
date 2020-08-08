package main

import "fmt"

func main() {
	fmt.Println(TestVar(5))
	fmt.Println(TestVar(18))
	fmt.Println(TestVar(20))
}

func TestVar(a int) int {
	var r int
	if a < 0 {
		return r
	}
	if a < 10 {
		r = a + 1
	} else {
		r := a + 2
		if r > 20 {
			r--
		}
	}
	return r
}

func TimeDiffDays(start int64, end int64) int {
	diff := int(end - start)
	if diff <= 0 {
		return 0
	}
	var days int
	if diff <= 86400 {
		days = 1
	} else {
		days := diff / 86400
		if diff%86400 > 0 {
			days += 1
		}
	}
	return days
}
