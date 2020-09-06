package main

import "fmt"

// -------------- 快速排序
// 有点难度

func QuickSortStep(arr []int, left int, right int) int {
	// 左起取第一个数作为基数， 每一趟走完 基数左边的数都小于基数，基数右边的都大于基数
	num := arr[left]
	for left < right {
		// 因为基数是从左起取第一个的， 所以要按照1，2顺序
		// 1, 先在右边找到第一个小于基数的数，拿走放在基数的位置上， 这时右边这个数是个待填充的坑
		for left < right && arr[right] >= num {
			right--
		}
		// 将基数右边小于基数的数给基数这个位置
		arr[left] = arr[right]
		// 2, 在基数左边查找大于基数的数， 拿走给上面右边的坑， 这时候这个数又成了一个填填充的坑
		for left < right && arr[left] <= num {
			left++
		}
		// 基数左边大于基数的，拿走给右边的坑
		arr[right] = arr[left]
	}
	// 把左边的坑填上， 这时候一趟走完，挖的坑都填上了。达到目的：基数左边的都小于基数，基数右边的都大于基数, 即确定了基数最终的位置
	arr[left] = num
	// 以基数为届， 由于左边没有排序，右边也没有排序，所以把两边的的集合再重复上面步骤，确定第二个基数最终的位置,以此类推
	return left
}

func QuickSort(arr []int, left int, right int) {
	if left < right {
		mid := QuickSortStep(arr, left, right)
		QuickSort(arr, left, mid-1)
		QuickSort(arr, mid+1, right)
	}
}

func main() {
	arr := []int{45, 38, 66, 87, 90, 88, 10, 25, 45}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
