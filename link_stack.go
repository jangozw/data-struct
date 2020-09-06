package main

import (
	"errors"
	"fmt"
)

//-------------------- 链栈 ---------------------------------
// 一个栈底节点其next为空 + n个栈数据节点, 不限制长度
// 链栈的节点结构，一个data存节点数据 ，一个next存下一个节点的指针
// ----------------------------------------------------------
type LinkStack struct {
	Data interface{}
	Next *LinkStack
}

func (s *LinkStack) EmptyStack() bool {
	return s.Next == nil
}

func (s *LinkStack) PushStack(val interface{}) {
	node := &LinkStack{
		Data: val,
		Next: s.Next,
	}
	s.Next = node
}

func (s *LinkStack) PopStack() (interface{}, error) {
	if s.EmptyStack() {
		return nil, errors.New("stack is empty")
	}
	node := s.Next
	s.Next = node.Next
	return node.Data, nil
}

func InitLinkStack() *LinkStack {
	return &LinkStack{
		Data: nil,
		Next: nil,
	}
}

func main() {
	stk := InitLinkStack()
	fmt.Println("入栈 1,2,3")
	stk.PushStack(1)
	stk.PushStack(2)
	stk.PushStack(3)

	fmt.Println("出栈:")
	fmt.Println(stk.PopStack())
	fmt.Println(stk.PopStack())
	fmt.Println(stk.PopStack())
	fmt.Println(stk.PopStack())
}
