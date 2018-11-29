package collections

import "GoStructAlgorithm/queue/base"

// 队列

type Queue struct {
    List *List
}

func (queue *Queue)Init()  {
    list := new(List)
    (*queue).List = list

    list.Init()
}

// 添加队列
func (queue *Queue)Enqueue(data base.Object)bool  {
    return (*queue).List.Append(data)
}

// 删除队列
func (queue *Queue)Dequeue()base.Object  {
    return(*queue).List.RemoveAt(0)
}

// 获取队列长度
func (queue *Queue)GetSize()uint64  {
    return (*queue).List.GetSize()
}

// 查看数据
func (queue *Queue)Peek()base.Object  {
    return (*queue).List.First()
}