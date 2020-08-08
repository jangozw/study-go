package main

import (
	"fmt"
)

func main() {
	// fmt.Println("=", 3/2)
	middleSearch()
	maopao()
}

func maopao() {
	// var new [6]int
	var count int
	arr := [6]int{1, 3, 2, 5, 7, 4}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
			count++
		}
	}
	fmt.Println(arr, count)
	// fmt.Print(arr)
}

func middleSearch() {
	// 二分查找前提要排序好的数据
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 50}
	var count int
	left := 0
	right := len(arr) // 为什么下表不是 len()-1, 因为除法导致(2+3)/2=2舍弃了小数, 这样最右边永远取不到下标3
	m, ok := digui(left, right, &count, arr, 50)
	if !ok {
		fmt.Println("couldn't find! count=", count)
	} else {
		fmt.Println("find index:", m, " count=", count)
	}
}

func digui(s, e int, count *int, arr [9]int, find int) (int, bool) {
	if *count > len(arr) {
		return 0, false
	}
	*count++
	m := (s + e) / 2
	if arr[m] == find {
		return m, true
	} else if arr[m] < find {
		return digui(m, e, count, arr, find)
	} else if arr[m] > find {
		return digui(s, m, count, arr, find)
	}
	return 0, false
}
