package main

import (
	"fmt"
)

// --------- 直接插入排序-------------------

func StraightInsertSorting(keys []int) []int {
	list := make([]int, len(keys)+1)
	for i := range keys {
		list[i+1] = keys[i]
	}
	for i := 2; i <= len(list)-1; i++ {
		j := i - 1
		list[0] = list[i]
		for list[0] < list[j] { // 后面数字 小于 前面
			list[j+1] = list[j]
			j-- // j 只会自减一次, 并且使循环自动跳出
		}
		list[j+1] = list[0]
	}
	return list[1:]
}

func main() {
	keys := []int{3, 4, 1, 45, 23, 12, 22}
	keys = StraightInsertSorting(keys)
	fmt.Println(keys)
}
