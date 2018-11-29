package collections

import "GoStructAlgorithm/queue/base"

// 栈

type Stack struct {
    List *List
}

func (stack *Stack)Init()  {
    list := new(List)
    (*stack).List = list
    list.Init()
}

// 插入栈数据
func (stack *Stack)Push(data base.Object) bool  {
    return (*stack).List.InsertAtHead(data)
}

// 移除栈
func (stack *Stack)Pop()bool  {
    list := (*stack).List
    data := list.RemoveAt(0)
    return data != nil
    //return (*stack).List.RemoveAt(0)
}

// 查看栈数据
func (stack *Stack)Peek() base.Object  {
    return (*stack).List.First()
}

// 获取栈数量
func (stack *Stack)GetSize() uint64 {
    return (*stack).List.GetSize()
}