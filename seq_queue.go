package main

import (
	"errors"
	"fmt"
)

// --------------------------------------------------------------------------------
// 顺序队列
// 缺点：会产生假溢出，不推荐使用，应该设计成循环队列，做这个例子是为了明白为什么要设计循环队列
// --------------------------------------------------------------------------------

const queMaxsize = 5

type SeqQue struct {
	Data  [queMaxsize]int // 队列元素 的数组
	Front int             // 队列首指针(数组下标 0开始)
	Rear  int             // 队列尾指针 (最大= max-1)
}

func (sq *SeqQue) EmptyQueue() bool {
	return sq.Front == sq.Rear
}

func (sq *SeqQue) EnQueue(val int) error {
	if sq.Rear >= queMaxsize-1 {
		fmt.Println("队列已满! 添加：", val)
		return errors.New("队列已满 ")
	}
	sq.Rear++
	sq.Data[sq.Rear] = val // 下标是从1 开始的
	return nil
}

func (sq *SeqQue) OutQueue() (int, error) {
	if sq.EmptyQueue() {
		return 0, errors.New("queue is empty")
	}
	if sq.Front+1 > queMaxsize-1 {
		fmt.Println("队列指针溢出")
		return 0, errors.New("队列指针溢出")
	}
	sq.Front++
	return sq.Data[sq.Front], nil
}

func main() {
	q := SeqQue{
		Data:  [queMaxsize]int{},
		Front: -1,
		Rear:  -1,
	}
	q.EnQueue(10)
	q.EnQueue(92)
	q.EnQueue(38)
	q.EnQueue(32)
	q.EnQueue(37)
	// q.EnQueue(38)
	fmt.Println(q.Data)
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	fmt.Println(q.OutQueue())
	// 这样产生里假溢出, 新元素增加不进来了， 因为sq.Front 到了最大值
	// 解决方法 ： 使用循环队列
	fmt.Println(q.EnQueue(44))
	fmt.Println(q.OutQueue())
}
