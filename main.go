package main

import (
    "fmt"
)

func main()  {
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
    array := GenerateRandomNumber(0, 100, 20)
    fmt.Println("排序前：", array)
    //start := time.Now().Unix()

    //array = BubbleSort(array)
    //array = SelectSort(array)
    array = InsertSort(array)
    //array = QuickSort(array,0,len(array)-1)
    //fmt.Printf("time = %d\n",time.Now().Unix()-start)
    fmt.Println("排序后：", array)
}
