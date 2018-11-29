package main

import (
    "GoStructAlgorithm/stack"
    "fmt"
    "GoStructAlgorithm/queue/base"
    "GoStructAlgorithm/queue/collections"
)

func main() {
    //rand.Seed(time.Now().UnixNano())
    //
    //for i :=0;i<100 ;i++  {
    //    fmt.Println(rand.Intn(10))
    //}

    //for i := 0; i < 100; i++ {
    //    result, _ := rand.Int(rand.Reader, big.NewInt(100))
    //    fmt.Println(result)
    //}
    //
    /*array := GenerateRandomNumber(0, 100, 20)
      fmt.Println("排序前：", array)
      //start := time.Now().Unix()
      stack.Stack{}

      //array = BubbleSort(array)
      //array = SelectSort(array)
      array = InsertSort(array)
      //array = QuickSort(array,0,len(array)-1)
      //fmt.Printf("time = %d\n",time.Now().Unix()-start)
      fmt.Println("排序后：", array)*/
    //TestStack_Push()
    testQueue()
}
func TestStack_Push() {
    var mStack stack.Stack
    mStack.Push(3)
    mStack.Push("test")
    mStack.Push("123")
    mStack.Push(false)
    fmt.Println(mStack.Len())
    mStack.Pop()
    fmt.Println(mStack.Len())

    if mStack.Len() == 1 {
        //t.Log("Pass Stack.Push")
    } else {
        //t.Error("Failed Stack.Push")
    }
}

type MyStr struct {
    name string
}

func match(data1, data2 base.Object) int {
    myStr1 := data1.(*MyStr)
    myStr2 := data2.(*MyStr)

    if (*myStr1).name == (*myStr2).name {
        return 0
    } else {
        return 1
    }
}

func testQueue()  {
    /*
    var lst List
    var plst *List = &lst
    plst.Init(match)

    myStr1 := new(MyStr)
    myStr2 := new(MyStr)
    myStr3 := new(MyStr)
    myStr4 := new(MyStr)
    myStrTag := new(MyStr)
    (*myStr1).name = "1"
    (*myStr2).name = "2"
    (*myStr3).name = "3"
    (*myStr4).name = "4"
    (*myStrTag).name = "2"
    plst.Append(myStr1)
    plst.Append(myStr2)
    plst.Append(myStr3)
    plst.Append(myStr4)
    fmt.Println(plst.GetSize())
    txt1 := lst.Next(myStrTag)
    txt2,r := txt1.(*MyStr)
    if (!r) {
        fmt.Println(r)
    } else {
        fmt.Println((*txt2).name)
    }
    lst := new (Queue)
    lst.Init()*/
    /*
    a := 1
    b := 2
    c := 3
    */

    //var d int
    /*lst.Enqueue("a")
    lst.Enqueue("c")
    lst.Enqueue("d")

    for ; lst.GetSize() > 0;  {
        fmt.Println(lst.Dequeue().(string))
    }
*/

    /*
        node := lst.GetHead().GetNext()
        fmt.Println(lst.GetSize())
        d = lst.Remove(node).(int)
        fmt.Println(lst.GetSize())
        fmt.Println(d)
        */

    set := new(collections.Set)
    set.Init(match)

    set1 := new(collections.Set)
    set1.Init(match)

    myStr1 := new(MyStr)
    myStr2 := new(MyStr)
    myStr3 := new(MyStr)
    myStr4 := new(MyStr)

    (*myStr1).name = "1"
    (*myStr2).name = "2"
    (*myStr3).name = "3"
    (*myStr4).name = "4"

    set.Insert(myStr1)
    set.Insert(myStr2)
    set.Insert(myStr3)
    set.Insert(myStr4)

    subSet := set1.IsSubSet(set)

    fmt.Println(subSet)

    myStr11 := new(MyStr)
    myStr21 := new(MyStr)
    myStr31 := new(MyStr)
    myStr41 := new(MyStr)

    (*myStr11).name = "1"
    (*myStr21).name = "2"
    (*myStr31).name = "3"
    (*myStr41).name = "4"

    set1.Insert(myStr11)
    set1.Insert(myStr21)
    set1.Insert(myStr31)
    set1.Insert(myStr41)

    set.Remove(myStr1)

    nset := set.InterSection(set1)

    var iterator collections.Iterator
    for iterator = nset.GetIterator(); iterator.HashNext(); {
        txt1 := iterator.Next()
        txt2,r := txt1.(*MyStr)
        if (!r) {
            fmt.Println(r)
        } else {
            fmt.Println((*txt2).name)
        }
    }
}