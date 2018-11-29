package collections

import (
    "GoStructAlgorithm/queue/base"
)

// 节点
type CNode struct {
    Data base.Object
    Next *CNode
}

// 单身链表
type CList struct {
    Size uint64
    Head *CNode
}

func (cList *CList) Init() {
    list := *cList
    list.Size = 0
    list.Head = nil
}

// 添加数据
func (cList *CList) Append(data base.Object) bool {
    node := new(CNode)
    (*node).Data = data

    if cList.GetSize() == 0 {
        (*cList).Head = node
    } else {
        item := cList.GetHead()
        // 遍历数据
        for ; (*item).Next != cList.GetHead(); item = (*item).Next {
        }
        (*item).Next = node
    }

    (*node).Next = (*cList).Head
    (*cList).Size++

    return true
}

// 插入下一个数据
func (cList *CList) InsertNext(elmt *CNode, data base.Object) bool {
    if elmt == nil {
        return false
    }
    node := new(CNode)
    (*node).Data = data
    (*node).Next = (*elmt).Next
    (*elmt).Next = node

    (*cList).Size++
    return true
}

// 移除数据
func (cList *CList) Remove(elmt *CNode) base.Object {
    if elmt == nil {
        return false
    }

    item := cList.GetHead()
    // 遍历数据
    for ; (*item).Next != elmt; item = (*item).Next {
    }

    (*item).Next = (*elmt).Next
    (*cList).Size --
    return elmt.GetData()

}

func (cList *CList) GetSize() uint64 {
    return (*cList).Size
}

func (cList *CList) GetHead() *CNode {
    return (*cList).Head
}

func (node *CNode) GetData() base.Object {
    return (*node).Data
}

func (node *CNode) GetNext() *CNode {
    return (*node).Next
}
