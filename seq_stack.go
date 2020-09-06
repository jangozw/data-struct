package main

import (
	"errors"
	"fmt"
)

// ------------------ 顺序栈 -------------------------
// 可以限制容量
// 实例化一个长度maxsize 数组，下标0作为栈底不存栈节点数据，用来判断是否空栈
// 可存储的数据元素为 maxsize - 1 个, 设置top作为数组栈顶元素的下标
// 压入栈： top++
// 出栈: 取数组中下标是top的数据元素然后另 top--
type SeqStack struct {
	Data    []int
	Top     int
	maxsize int
}

func (s *SeqStack) EmptyStack() bool {
	return s.Top == 0
}

func (s *SeqStack) FullStack() bool {
	return s.Top == s.maxsize-1
}

func (s *SeqStack) PushStack(val int) error {
	if s.FullStack() {
		return errors.New("stack is full")
	}
	s.Top++
	s.Data[s.Top] = val
	return nil
}

func (s *SeqStack) PopStack() (int, error) {
	if s.EmptyStack() {
		return 0, errors.New("stack is empty")
	}
	val := s.Data[s.Top]
	s.Top--
	return val, nil
}

// 0 作为栈底，用来判断是否空栈， 数组实际需要len+1个容量空间
func InitSeqStack(len int) *SeqStack {
	return &SeqStack{
		Data:    make([]int, len+1),
		Top:     0,
		maxsize: len + 1,
	}
}

func main() {
	fmt.Println("初始化容量为3的栈")
	stack := InitSeqStack(3)

	fmt.Println("入栈 1", stack.PushStack(1))
	fmt.Println("入栈 2", stack.PushStack(2))
	fmt.Println("入栈 3", stack.PushStack(3))
	fmt.Println("入栈 4", stack.PushStack(4))

	fmt.Println("出栈: ")
	fmt.Println(stack.PopStack())
	fmt.Println(stack.PopStack())
	fmt.Println(stack.PopStack())
	fmt.Println(stack.PopStack())
}
