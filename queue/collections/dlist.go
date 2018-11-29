package collections

import (
    "GoStructAlgorithm/queue/base"
)

// 节点
type DNode struct {
    Data base.Object
    Prev *DNode
    Next *DNode
}

// 双端链表
type DList struct {
    Size uint64
    Head *DNode // 头部
    Tail *DNode // 尾部
}

// 双链表初始化
func (dList *DList) Init() {
    list := *(dList)
    list.Size = 0
    list.Head = nil
    list.Tail = nil
}

func (dList *DList) GetSize() uint64 {
    return (*dList).Size
}

func (dList *DList) GetHead() *DNode {
    return (*dList).Head
}

// 获取尾部节点
func (dList *DList) GetTail() *DNode {
    return (*dList).Tail
}

// 判断是否是头部
func (dList *DList) IsHead(elmt *DNode) bool {
    return dList.GetHead() == elmt
}

// 判断是否是尾部
func (dList *DList) IsTail(elmt *DNode) bool {
    return dList.GetTail() == elmt
}

// 插入数据
func (dList *DList) Append(data base.Object) {
    newNode := new(DNode)
    (*newNode).Data = data

    if (*dList).GetSize() == 0 {
        (*dList).Head = newNode
        (*dList).Tail = newNode
        (*newNode).Prev = nil
        (*newNode).Next = nil
    } else {
        (*newNode).Prev = (*dList).Tail
        (*newNode).Next = nil
        (*((*dList).Tail)).Next = newNode
        (*dList).Tail = newNode
    }
    (*dList).Size++

}

// 插入下一个数据
func (dList *DList) InsertNext(elmt *DNode, data base.Object) bool {

    if elmt == nil {
        return false
    }

    if dList.IsTail(elmt) {
        dList.Append(data)
    } else {
        newNode := new(DNode)
        (*newNode).Data = data
        (*newNode).Prev = elmt
        (*newNode).Next = (*elmt).Next

        (*elmt).Next = newNode
        (*((*newNode).Next)).Prev = newNode
        (*dList).Size++
    }
    return true

}

// 往头部插入数据
func (dList *DList) InsertPrev(elmt *DNode, data base.Object) bool {
    if elmt == nil {
        return false
    }

    if dList.IsHead(elmt) {
        newNode := new(DNode)
        (*newNode).Data = data
        (*newNode).Next = dList.GetHead()
        (*newNode).Prev = nil

        (*((*dList).Head)).Prev = newNode

        dList.Head = newNode
        dList.Size++
        return true
    } else {
        prev := (*elmt).Prev
        return dList.InsertNext(prev, data)
    }
}

func (dList *DList) Remove(elmt *DNode) base.Object {
    if elmt == nil {
         return nil
    }

    prev := (*elmt).Prev
    next := (*elmt).Next

    if dList.IsHead(elmt) {
        dList.Head = next
    }else {
        (*prev).Next = next
    }

    if dList.IsTail(elmt) {
        dList.Tail = prev
    } else {
        (*next).Prev = prev
    }

    dList.Size++
    return (*elmt).Data
}

// 查找节点，参数1：查找的数据，参数2：执行方法切片。返回所查找到的节点
func (dList *DList)Search(data base.Object,yourMatch ...base.MatchFun) *DNode {
    if dList.GetSize() == 0 {
        return nil
    }
    // 给方法取个别名
    match := DefaultMatch
    // 判断方法切片
    if len(yourMatch) >0{
        match = yourMatch[0]
    }

    node := dList.GetHead()
    // 遍历节点
    for ;node != nil ;node.GetNext()  {
        // 如果节点的数据跟传入的数据相同就返回
        if match(node.GetData(),data) == 0 {
            break
        }
    }
    return node
}


func (dNoce *DNode) GetData() base.Object {
    return (*dNoce).Data
}

func (dNoce *DNode) GetNext() base.Object {
    return (*dNoce).Next
}

func (dNoce *DNode) GetPrev() base.Object {
    return (*dNoce).Prev
}
