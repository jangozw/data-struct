package main

//--------------------- 链队列 ------------------------------
// 长度不限
// 原理： 使用一个带头节点的单链表表示队列，长度不限制, 头指针指向链表的头节点， 头节点的next指针指向队列的首节点,
// 头节点用来获取队列首节点和判断队列是否为空
// 队列的尾指针指向队列的尾巴节点，即但链表最后一个节点，为节点的next指向为空
// 为空队列的条件：队列的头指针 == 尾指针

import (
	"errors"
	"fmt"
)

// 队列里的节点结构
// 每个节点包含一个data域和下一个节点的指针next， next就是下一个节点
type LinkQueueNode struct {
	Data int
	Next *LinkQueueNode
}

// 队列结构体， 包含一个头指针，尾指针
type LinkQueue struct {
	Front *LinkQueueNode
	Rear  *LinkQueueNode
}

// 判断队列为空
func (q *LinkQueue) EmptyQueue() bool {
	return q.Front == q.Rear
}

// 入队列
// 将一个节点插入队列， 从尾入，这个新增节点就成了尾指针
// 要修改：1，当前的尾节点（next为空）的next需要指向为新加入的节点， 2，将新节点设置为队列尾巴指针
func (q *LinkQueue) EnQueue(val int) {
	node := &LinkQueueNode{
		Data: val,
		Next: nil, // 新加的节点必然是尾节点所以next为空
	}
	// 旧的尾节点退位了，让给新加入的节点
	q.Rear.Next = node
	// 新加入的节点担任尾节点
	q.Rear = node
}

// 出队列
// 1, 使用队列头节点的next获取队列首节点
// 2，修改头节点的next指针使其指向首节点的下一个节点， 如果首节点的下一个节点是尾为节点

func (q *LinkQueue) OutQueue() (int, error) {
	if q.EmptyQueue() {
		return 0, errors.New("queue is empty")
	}
	// 要取出的节点
	node := q.Front.Next
	// 队列头指针需要修改
	q.Front.Next = node.Next
	// 如果取的是尾节点，那取完后要将队列置空即尾节点等于头节点回到初始化状态
	if node.Next == nil {
		q.Rear = q.Front
	}
	return node.Data, nil
}

// 访问队列首节点， 注意不是出队列
func (q *LinkQueue) GetHead() int {
	if q.Front.Next == nil {
		return 0
	}
	return q.Front.Next.Data
}

// 初始化一个链队列
func InitLinkQueue() *LinkQueue {
	// 队列头节点
	headNode := &LinkQueueNode{
		Data: 0,
		Next: nil,
	}
	return &LinkQueue{
		Front: headNode,
		Rear:  headNode,
	}
}

func main() {
	q := InitLinkQueue()

	fmt.Println("入队列:1,2,3")
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	fmt.Println("出队列结果:")
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
}
