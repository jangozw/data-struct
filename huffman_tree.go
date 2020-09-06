package main

import (
	"encoding/json"
	"fmt"
)

// Huffman tree node
type HuffmanNode struct {
	Weight float32
	Parent int // 双亲节点数组下标
	Left   int // 左孩子节点下标
	Right  int // 右孩子节点下标
}

/**
生成哈夫曼树的要点：
1， n个节点 要生成哈夫曼树需要新增 n-1 个节点，总共2n-1个
2，每次查找最小权重的两个节点，合并成一个新节点， 作为原来节点的双亲，新节点的左右孩子为 该两个被合并的节点
*/

func main() {
	// 权重节点
	weight := []float32{0.2, 0.3, 0.25, 0.25}
	// 设置一个超过最大权重的值
	var weightMaxLimit float32 = 10000
	tree := makeHuffmanTree(weight, weightMaxLimit)
	fmt.Println(ToJson(tree))
}

func makeHuffmanTree(weight []float32, weightMaxLimit float32) []HuffmanNode {
	num := len(weight)
	// 将生成的哈夫曼树的总结的数 （根据哈夫曼树的性质）
	newNodeNum := num - 1
	// 根据哈夫曼树性质，num 个树转化成哈夫曼树将增加 n-1 个节点，总结点为n-1+n = 2n-1个
	allNodeNum := num + newNodeNum
	// 初始化哈夫曼树
	// var huffman [7]HuffmanNode
	huffman := make([]HuffmanNode, allNodeNum)
	for i := 0; i < allNodeNum; i++ {
		huffman[i] = HuffmanNode{
			Weight: 0,
			Parent: -1,
			Left:   -1,
			Right:  -1,
		}
		if i < num {
			huffman[i].Weight = weight[i]
		}
	}
	// 新增加 的num -1 个节点，上面已经初始化了，现在把他们填上具体的值
	for i := 0; i < num-1; i++ {
		// 所有非空节点中找出最小的两个权重值 相加 = 新节点的权重， 再修改原节点的 双亲节点指针指向 新节点
		min1, min2 := weightMaxLimit, weightMaxLimit
		var min1Index, min2Index int
		for j := 0; j < num+i; j++ {
			// 只能再没有根的节点里比出最小的两位
			if huffman[j].Weight < min1 && huffman[j].Parent == -1 {
				min2 = min1           // 获得最小两个值的关键点
				min2Index = min1Index // 获得最小两个值的关键点
				min1 = huffman[j].Weight
				min1Index = j
			} else if huffman[j].Weight < min2 && huffman[j].Parent == -1 {
				min2 = huffman[j].Weight
				min2Index = j
			}
		}
		// 将权重最小的两个节点合并，每次合并得到一个新节点 作为根
		// 合并成的新节点
		newNodeIndex := num + i
		huffman[newNodeIndex].Weight = min1 + min2
		huffman[newNodeIndex].Left = min1Index
		huffman[newNodeIndex].Right = min2Index
		// 修改被合并节点的双亲节点指针
		huffman[min1Index].Parent = newNodeIndex
		huffman[min2Index].Parent = newNodeIndex
	}
	return huffman
}

func ToJson(v interface{}) string {
	jb, _ := json.Marshal(v)
	return string(jb)
}
