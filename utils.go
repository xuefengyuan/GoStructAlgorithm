package main

import (
    "math/rand"
    "time"
)

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

