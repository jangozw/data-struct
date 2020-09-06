package main

import (
	"encoding/json"
	"fmt"
)

// --------------二叉排序树的查找和构建------------------

type BstNode struct {
	Key    int
	LChild *BstNode
	RChild *BstNode
}

// parentNode 会被改变
func BstSearch(node *BstNode, key int, parentNode *BstNode) (*BstNode, *BstNode) {
	if node == nil {
		fmt.Println("parent:", parentNode.Key, "curr")
		return nil, parentNode
	} else {
		if node.Key == key {
			return node, parentNode
		} else if key < node.Key {
			return BstSearch(node.LChild, key, node)
		} else {
			return BstSearch(node.RChild, key, node)
		}
	}
}

func BstInsertNode(node *BstNode, key int) {
	findNode, parentNode := BstSearch(node, key, nil)
	if findNode != nil {
		return
	}
	// 根据上面的递归查询, parentNode 必然不为空
	if parentNode != nil {
		addNode := &BstNode{
			Key:    key,
			LChild: nil,
			RChild: nil,
		}
		if key < parentNode.Key {
			parentNode.LChild = addNode
		} else {
			parentNode.RChild = addNode
		}
	}
}

func main() {
	// 先初始化根节点
	rootNode := &BstNode{
		Key:    19,
		LChild: nil,
		RChild: nil,
	}
	keys := []int{19, 14, 22, 1, 66, 21, 83, 27, 56, 13, 10}
	for _, v := range keys {
		BstInsertNode(rootNode, v)
	}
	fmt.Println(ToJson1(rootNode))
}

func ToJson1(v interface{}) string {
	jb, _ := json.Marshal(v)
	return string(jb)
}
