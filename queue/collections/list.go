package collections

import (
    "GoStructAlgorithm/queue/base"
)

// 节点
type Node struct {
    Data base.Object
    Next *Node
}

// 双向链表
type List struct {
    Size    uint64
    Head    *Node
    Tail    *Node
    MyMatch base.MatchFun
}

// 默认的数据比较方法
func DefaultMatch(data1, data2 base.Object) int {
    if data1 == data2 {
        return 0
    } else {
        return 1
    }
}

// 数据比较方法
func (list *List) Match(data1, data2 base.Object) int {
    var match base.MatchFun = nil
    if (*list).MyMatch == nil {
        match = DefaultMatch
    } else {
        match = (*list).MyMatch
    }
    return match(data1, data2)
}

// 创建节点
func (list *List) CreateNode(data base.Object) *Node {
    node := new(Node)
    (*node).Data = data
    (*node).Next = nil
    return node
}

// 获取下一个节点
func GetNextNode(node *Node) *Node {
    return (*node).Next
}

// 获取头部节点
func (list *List) GetHead() *Node {
    return (*list).Head
}

// 获取尾部节点
func (list *List) GetTail() *Node {
    return (*list).Tail
}

func (list *List) GetSize() uint64 {
    return (*list).Size
}

func (list *List) IsEmpty() bool {
    return (*list).GetSize() == 0
}

// 获取头部数据
func (list *List) First() base.Object {
    if list.GetSize() == 0 {
        return nil
    } else {
        return (*(list.GetHead())).Data
    }
}

// 获取尾部数据
func (list *List) Last() base.Object {
    if list.GetSize() == 0 {
        return nil
    } else {
        return (*(list.GetTail())).Data
    }
}

// 获取下一个数据
func (list *List) Next() base.Object {
    head := list.GetHead()

    for i := head; i != nil; i = GetNextNode(i) {
        next := GetNextNode(i)
        if next == nil {
            return nil
        } else {
            return next.GetData()
        }
    }
    return nil
}

func (node *Node) GetData() base.Object {
    return (*node).Data
}

// 往后面插入一个节点
func (list *List) InsertAfterNode(node *Node, data base.Object) bool {
    // TODO 未实现
    return true
}

// 根据索引删除节点
func (list *List) RemoveAt(index uint64) base.Object {
    size := list.GetSize()

    if index >= size {
        return nil
    } else if size == 1 { // 只有一个
        // 得到头部的节点
        node := list.GetHead()
        (*list).Head = nil
        (*list).Tail = nil
        (*list).Size = 0
        return (*node).Data
    } else if index == 0 { // 移除头部
        // 得到头部的节点
        node := list.GetHead()
        (*list).Head = (*node).Next
        (*list).Size--
        return (*node).Data
    } else if index == (size - 1) { // 移除尾部
        // 得到头部的节点
        preNode := list.GetHead()
        // 从头部遍历到尾部
        for i := uint64(2); i < size; i++ {
            preNode = (*preNode).Next
        }
        // 得到尾部的节点
        tail := list.GetTail()
        (*list).Tail = preNode
        preNode.Next = nil
        (*list).Size--
        return (*tail).Data
    } else { // 正常删除数据
        // 得到头部的节点
        preNode := list.GetHead()
        // 从头部遍历到尾部
        for i := uint64(2); i < size; i++ {
            preNode = (*preNode).Next
        }
        node := (*preNode).Next
        nextNode := (*node).Next
        (*node).Next = nextNode
        (*list).Size--

        return (*node).Data
    }
}

// 根据数据删除节点
func (list *List) Remove(data base.Object) bool {
    if (data == nil || list.IsEmpty()) {
        return false
    }
    head := list.GetHead()

    if (list.Match(head.GetData(), data) == 0) {
        (*list).Head = GetNextNode(head)
    } else {
        cur := head
        // 得到头部的下一下节点
        next := GetNextNode(head)
        // 不断的遍历下一个节点
        for ; next != nil; next = GetNextNode(next) {
            if list.Match(data, next.GetData()) == 0 {
                (*cur).Next = GetNextNode(next)
                break
            }
            cur = next
        }
        if next == nil {
            return false
        }
    }
    (*list).Size--
    return false
}

// 判断数据是否在链节点中
func (list *List) IsMember(data base.Object) bool {
    if list.IsEmpty() {
        return false
    }

    head := list.GetHead()
    for i := head; i != nil; i = GetNextNode(i) {
        // 如果数据相同，则表示在
        if list.Match(data, i.GetData()) == 0 {
            return true
        }
    }

    return false
}

// List初始化
func (list *List) Init(yourMatch ...base.MatchFun) {
    (*list).Size = 0
    (*list).Head = nil
    (*list).Tail = nil
    if len(yourMatch) == 0 {
        (*list).MyMatch = nil
    } else {
        (*list).MyMatch = yourMatch[0]
    }
}

// 正常添加数据
func (list *List) Append(data base.Object) bool {
    newItem := new(Node)
    (*newItem).Data = data
    (*newItem).Next = nil
    if (*list).Size == 0 {
        (*list).Head = newItem
        (*list).Tail = (*list).Head
    } else {
        oldNode := (*list).Tail
        (*oldNode).Next = newItem
        (*list).Tail = newItem
    }

    (*list).Size++

    return true
}

// 往头部插入数据
func (list *List) InsertAtHead(data base.Object) bool {
    newNode := list.CreateNode(data)

    (*newNode).Next = list.GetHead()
    list.Head = newNode
    (*list).Size++

    return true
}

// 根据索引获取数据
func (list *List) GetAt(index uint64) base.Object {
    size := list.GetSize()
    if index > size {
        return nil
    } else if index == 0{
        return list.First()
    } else if index == (size -1){
        return list.Last()
    } else {
        item := list.GetHead()
        for i := uint64(2); i<size ; i++  {
            if i == index{
                break
            }
            item = (*item).Next
        }
        return item.GetData()
    }

}

// 根据下标插入数据
func (list *List)insertAt(index uint64,data base.Object) bool  {
    size := list.GetSize()

    if index > size {
        return false
    } else if index == size{ // 正常添加数据
        return list.Append(data)
    } else if index == 0{ // 添加到头部
        return list.InsertAtHead(data)
    } else {
        newNode := list.CreateNode(data)
        prevIndex := index -1
        previtem := list.GetHead()
        for i := uint64(0);i<size ; i++ {
            if i == prevIndex {
                break
            }
            previtem = (*previtem).Next

        }

        (*newNode).Next = (*previtem).Next
        (*previtem).Next = newNode

        (*list).Size++

        return true
    }

}

// 清除数据
func (list *List) Clear() {
    list.Init()
}
