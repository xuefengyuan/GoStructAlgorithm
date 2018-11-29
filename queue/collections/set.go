package collections

import (
    "GoStructAlgorithm/queue/base"
)

type Set struct {
    List *List
}

type SetIterator struct {
    Index uint64
    Set   *Set
}

func (set *Set) Remove(data base.Object) bool {
    return (*set).List.Remove(data)
}

func (set *Set) IsEmpty() bool {
    return (*set).List.IsEmpty()
}

// 判断数据是否存在
func (set *Set) IsMember(data base.Object) bool {
    return (*set).List.IsMember(data);
}

// 根据索引获取数据
func (set *Set) GetAt(index uint64) base.Object {
    return (*set).List.GetAt(index)
}

func (set *Set) GetSize() uint64 {
    return (*set).List.GetSize()
}

// 初始化
func (set *Set) Init(match ...base.MatchFun) {
    list := new(List)
    (*set).List = list

    if len(match) == 0 {
        list.Init()
    } else {
        list.Init(match[0])
    }
}

// 插入数据
func (set *Set) Insert(data base.Object) bool {
    if !set.IsMember(data) {
        return (*set).List.Append(data)
    }
    return false
}

func (set *Set) Union(set1 *Set) *Set {
    if set1 == nil {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).List)).MyMatch)

    if set.IsEmpty() && set1.IsEmpty() {
        return nSet
    }

    for i := uint64(0); i < set.GetSize(); i++ {
        nSet.Insert(set.GetAt(i))
    }

    var data base.Object
    for i := uint64(0); i < set1.GetSize(); i++ {
        data = set1.GetAt(i)
        if !nSet.IsMember(data) {
            nSet.Insert(data)
        }
    }
    return nSet
}

func (set *Set) InterSection(set1 *Set) *Set {
    if set1 == nil {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).List)).MyMatch)
    if set.IsEmpty() && set1.IsEmpty() {
        return nSet
    }

    fSet := set
    sSet := set1
    lenth := set.GetSize()

    if set1.GetSize() < lenth {
        fSet = set1
        sSet = set
    }

    var data base.Object
    for i := uint64(0); i < lenth; i++ {
        data = fSet.GetAt(i)
        if sSet.IsMember(data) {
            nSet.Insert(data)
        }
    }

    return nSet
}

// 差异
func (set *Set) Difference(set1 *Set) *Set {
    if set1 == nil {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).List)).MyMatch)
    if set.IsEmpty() {
        return nSet
    }

    var data base.Object
    for i := uint64(0); i < set.GetSize(); i++ {
        data = set.GetAt(i)
        if !set1.IsMember(data) {
            nSet.Insert(data)
        }
    }
    return nSet
}

func (set *Set) IsSubSet(subSet *Set) bool {
    if set == nil {
        return false
    }

    if subSet == nil {
        return true
    }

    for i := uint64(0); i < subSet.GetSize(); i++ {
        if !set.IsMember(subSet.GetAt(i)) {
            return false
        }
    }
    return true
}

func (set *Set) Equals(set1 *Set) bool {
    if set == nil || set1 == nil {
        return false
    }

    if set.IsEmpty() || set1.IsEmpty() {
        return false
    }

    nSet := set.InterSection(set1)

    return set.GetSize() == nSet.GetSize()
}

func (set *Set) GetIterator() *SetIterator {
    iterator := new(SetIterator)
    (*iterator).Index = 0
    (*iterator).Set = set

    return iterator
}

func (iterator *SetIterator) HashNext() bool {
    set := (*iterator).Set
    index := (*iterator).Index
    return index < set.GetSize()
}

func (iterator *SetIterator) Next() base.Object {
    set := (*iterator).Set
    index := (*iterator).Index
    if index < set.GetSize() {
        data := set.GetAt(index)
        (*iterator).Index++
        return data
    }
    return nil
}
