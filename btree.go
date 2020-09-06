// 二叉树的遍历
package main

import "fmt"

type BtNode struct {
	Data                   string
	LChild, RChild, Parent *BtNode
}

// var Btree = [6]int{'A', 'B', 'C', 'D', 'E', 'F'}
var Bt *BtNode

func init() {
	Bt = &BtNode{
		Data: "A",
		LChild: &BtNode{
			Data: "B",
			LChild: &BtNode{
				Data:   "D",
				LChild: nil,
				RChild: nil,
			},
			RChild: &BtNode{
				Data: "E",
				LChild: &BtNode{
					Data:   "G",
					LChild: nil,
					RChild: nil,
				},
				RChild: nil,
			},
		},
		RChild: &BtNode{
			Data:   "C",
			LChild: nil,
			RChild: &BtNode{
				Data:   "F",
				LChild: nil,
				RChild: nil,
			},
		},
	}
}

func main() {
	fmt.Print("先序遍历：")
	PreOrder(Bt)
	fmt.Println()

	fmt.Print("中序遍历：")
	InOrder(Bt)
	fmt.Println()

	fmt.Print("后序遍历：")
	PostOrder(Bt)

	fmt.Println("\n二叉链表叶子节点个数:", leafNodeNum(Bt))
	fmt.Println("总结点个数", nodeNum(Bt))
}

func nodeNum(tree *BtNode) int {
	if tree == nil {
		return 0
	}
	return 1 + nodeNum(tree.LChild) + nodeNum(tree.RChild)
}

// 二叉链表叶子节点个数
func leafNodeNum(tree *BtNode) int {
	if tree == nil {
		return 0
	}
	if tree.RChild == nil && tree.LChild == nil {
		return 1
	}
	return leafNodeNum(tree.RChild) + leafNodeNum(tree.LChild)
}

// 中序便利二叉树
func PreOrder(n *BtNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Data, " ")
	PreOrder(n.LChild)
	PreOrder(n.RChild)
}

// 中序便利二叉树
func InOrder(n *BtNode) {
	if n == nil {
		return
	}
	InOrder(n.LChild)
	fmt.Print(n.Data, " ")
	InOrder(n.RChild)
}

// 中序便利二叉树
func PostOrder(n *BtNode) {
	if n == nil {
		return
	}
	PostOrder(n.LChild)
	PostOrder(n.RChild)
	fmt.Print(n.Data, " ")
}
