package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 2, 5}
	a[0], a[1] = a[1], a[0]
	fmt.Println(a)
}

func ju() {
	a := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = j + 1
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[j][i] = a[i][j]
		}
		fmt.Println(a[i])
	}
}
