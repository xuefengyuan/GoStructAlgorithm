package main

import (
    "fmt"
    "time"
    "math/rand"
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


    array := GenerateRandomNumber(0, 100, 20)

    fmt.Println("排序前：", array)
    //array = BubbleSort(array)
    array = SelectSort(array)
    fmt.Println("排序后：", array)
}

/*
  生成不重复的随机数

  起始数

  终止数

  生成的个数

*/
func GenerateRandomNumber(start int, end int, count int) []int {
    // 判断设置的个数是否合理
    if end < start || (end-start) < count {
        return nil
    }

    nums := make([]int, 0)
    // 创建随机数种子 进行数据混淆
    rand.Seed(time.Now().UnixNano())
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for len(nums) < count {
        num := r.Intn((end - start)) + start

        exist := false
        for _, v := range nums {
            if v == num {
                exist = true
                break
            }
        }

        if !exist {
            nums = append(nums, num)
        }
    }

    return nums
}

