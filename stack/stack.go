package stack

import "errors"

/*
  栈的操作
  栈是一组记录，表现形式为先进后出
*/
type Stack []interface{}

// 获取栈的数量
func (stack Stack) Len() int {
    return len(stack)
}

// 判断是否为空
func (stack Stack) IsEmpey() bool {
    return len(stack) == 0
}

// 获取栈的容量
func (stack Stack) Cap() int {
    return cap(stack)
}

// 添加数据
func (stack *Stack) Push(value interface{}) {
    *stack = append(*stack, value)
}

// 返回顶部
func (stack Stack) Top() (interface{}, error) {
    if stack.IsEmpey() {
        return nil, errors.New("Out of index, len is 0")
    }
    return stack[len(stack)-1], nil
}

// 删除
func (stack *Stack) Pop() (interface{}, error) {

    theStack := *stack
    if stack.IsEmpey() {
        return nil, errors.New("Out of index, len is 0")
    }
    value := theStack[len(theStack)-1]
    *stack = theStack[:len(theStack)-1]
    return value, nil
}
