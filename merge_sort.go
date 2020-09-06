package main

import "fmt"

// 归并排序
func main() {
	arr := []int{10, 18, 4, 3, 6, 12, 1, 9, 15, 8}
	temp := make([]int, len(arr))
	MsStateSort(arr, 0, len(arr)-1, temp)
	fmt.Println("sorted", arr)
}

func MsStateSort(arr []int, left int, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		MsStateSort(arr, left, mid, temp)
		MsStateSort(arr, mid+1, right, temp)

		fmt.Println(temp)
		MsStateMerge(arr, left, mid, right, temp)

	}
}

func MsStateMerge(arr []int, left int, mid int, right int, temp []int) {
	i := left
	k := mid + 1
	t := 0
	for i <= mid && k <= right {
		if arr[i] <= arr[k] {
			temp[t] = arr[i]
			i++
		} else {
			temp[t] = arr[k]
			k++
		}
		t++
	}
	for i <= mid {
		temp[t] = arr[i]
		i++
		t++
	}
	for k <= right {
		temp[t] = arr[k]
		k++
		t++
	}

	t = 0
	for left <= right {
		arr[left] = temp[t]
		t++
		left++
	}
	// fmt.Println(temp)
}
