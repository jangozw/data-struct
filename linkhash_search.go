package main

import "fmt"

//-------------------------------------链地址法的散列查找算法-------------
// 原理： 将同义词都放在某个散列值的线性表里， 根据散列值找到线性表，在遍历取得,以解决冲突
//
const maxHashTableSize = 13

//
type LinkHashNode struct {
	// 实际值
	Key int
	// 同义词链表
	Next *LinkHashNode
}

type LinkHashTable [maxHashTableSize]*LinkHashNode

func GetLinkHashValue(key int) int {
	return key % maxHashTableSize
}

func InitLinkHashTable() *LinkHashTable {
	var t LinkHashTable
	return &t
}

func (t *LinkHashTable) Get(key int) *LinkHashNode {
	hash := GetLinkHashValue(key)
	node := t[hash]
	for node != nil && node.Key != key {
		node = node.Next
	}
	return node
}

func (t *LinkHashTable) Set(key int) {
	find := t.Get(key)
	if find == nil {
		hash := GetLinkHashValue(key)
		addNode := &LinkHashNode{
			Key:  key,
			Next: t[hash],
		}
		t[hash] = addNode
	}
}

func main() {
	hashTable := InitLinkHashTable()
	keys := []int{26, 41, 25, 5, 7, 15, 12, 49, 51, 31, 62}
	for _, v := range keys {
		hashTable.Set(v)
	}

	// fmt.Println(hashTable)
	fmt.Println(hashTable.Get(25))
	fmt.Println(hashTable.Get(49))
}
