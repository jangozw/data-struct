package main

import "fmt"

// ---------------- 直接选择排序

// 直接选择排序根冒泡排序有啥区别？就是交换的过程写在内循环里面还是外面的区别

func SelectSort(keys []int) {
	for i := 0; i < len(keys); i++ {
		min := i
		for j := i + 1; j < len(keys); j++ {
			if keys[j] < keys[min] {
				min = j
			}
		}
		//-----交换写在内循环外面是选择排序， 写在内循环里面是冒泡排序
		if min != i {
			tmp := keys[min]
			keys[min] = keys[i]
			keys[i] = tmp
		}
		//--------------------------
	}
}

func main() {
	arr := []int{45, 38, 66, 87, 90, 88, 10, 25, 45}
	SelectSort(arr)
	fmt.Println(arr)
}
