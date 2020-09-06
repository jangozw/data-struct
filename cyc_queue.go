package main

import (
	"errors"
	"fmt"
)

/*
	----------------------------------------------------------------
	循环队列
	// 利用数组存储队列节点， 队列设置一个头尾指针作为数组的下标入队和出队
	原理：为克服假上溢，队首和队尾之间保留一个空元素即可认为队列已满。
	循环队列实际入队的节点 最多是 maxsize -1 个，因为还有空一个节点判断队列是否已满
	以最大5为例，因为 0 % 5 = 0 ， 5%5=0， 不好区分是否已满
	----------------------------------------------------------------
*/

type CycQueue struct {
	maxsize int
	Data    []int
	Front   int
	Rear    int
}

func (q *CycQueue) EmptyQueue() bool {
	return q.Front == q.Rear
}

func (q *CycQueue) EnQueue(val int) error {
	key := (q.Rear + 1) % q.maxsize
	if key == q.Front {
		return errors.New("Queue is full ")
	}
	q.Rear = key
	q.Data[q.Rear] = val
	return nil
}

func (q *CycQueue) OutQueue() (int, error) {
	if q.EmptyQueue() {
		return 0, errors.New("queue is empty")
	}
	q.Front = (q.Front + 1) % q.maxsize
	val := q.Data[q.Front]
	return val, nil
}

// 队列中数据节点的数量, 注意不包含头节点
// 数组做的循环队列中 如果front > rear 说明在 front 在左边， rear 在右
// 如5个容量的数组中 [0 1 2 3 4][ 0 1 2 3 4]，front =4， rear=3 时候数据节点个数=5-4+3 (front 永远是头节点下标不算在内，取数据时候从front+1开始)
func (q *CycQueue) DataNodeNum() int {
	if q.Front == q.Rear {
		return 0
	} else if q.Front < q.Rear {
		return q.Rear - q.Front
	} else {
		return q.maxsize - q.Front + q.Rear
	}
}

// 初始化队列， len 是实际存储节点的个数， len+1 是初始化队列内部的长度，要增加一个空闲节点判断溢出
func InitCycQueue(len int) *CycQueue {
	return &CycQueue{
		Data:    make([]int, len+1),
		Front:   0,
		Rear:    0,
		maxsize: len + 1,
	}
}

func main() {
	q := InitCycQueue(30)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)

	fmt.Println(q.Data)
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
}
