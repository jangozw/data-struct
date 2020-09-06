package main

import "fmt"

// ------------------------------ 静态表的顺序查找, 二分查找----------------------

// 数据元素结构体
type SeqTableElem struct {
	Key int
}

// 静态表
type SeqTable struct {
	Elem []SeqTableElem // 数据元素
	Last int            // 数据元素的最后的一个下标
}

// 顺序查找
func (s *SeqTable) SeqSearch(key int) int {
	s.Elem[0].Key = key
	i := s.Last
	for s.Elem[i].Key != key {
		i--
	}
	return i
}

// 二分查找，前提是保证存储的数据是按从小到大顺序一定的
func (s *SeqTable) BinSearch(key int) int {
	var c int
	var find int
	left := 1
	right := s.Last
	for left <= right {
		c++
		mid := (left + right) / 2
		midKey := s.Elem[mid].Key
		if midKey == key {
			find = mid
			break
		} else if midKey < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("binSearch count:", c)
	return find
}

// 初始化一个静态表
func InitSeqTable(maxsize int) *SeqTable {
	return &SeqTable{
		Elem: make([]SeqTableElem, maxsize+1),
		Last: maxsize - 1,
	}
}

func main() {
	t := InitSeqTable(9)
	t.Elem[1].Key = 10
	t.Elem[2].Key = 20
	t.Elem[3].Key = 30
	t.Elem[4].Key = 40
	t.Elem[5].Key = 50
	t.Elem[6].Key = 60
	t.Elem[7].Key = 70
	t.Elem[8].Key = 80
	t.Elem[9].Key = 90

	// position := t.SeqSearch(20)
	// fmt.Println("seqSearch find the key position: ", position)

	// position2 := t.BinSearch(10)
	// fmt.Println("binSearch find the key position: ", position2)
	t.BinSearch(10)
	t.BinSearch(20)
	t.BinSearch(30)
	t.BinSearch(40)
	t.BinSearch(50)
	t.BinSearch(60)
	t.BinSearch(70)
	t.BinSearch(80)
	t.BinSearch(90)
}
