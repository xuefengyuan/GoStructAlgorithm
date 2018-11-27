package main

/*
  选择排序
*/
func SelectSort(array []int) []int {
    for i := 0; i < len(array)-1; i++ { // 1.做第i趟排序
        var k = i
        for j := k + 1; j < len(array); j++ { // 2.遍历查找最小的记录
            if array[j] < array[k] {
                // 3.记下目前找到的最小值所在的位置
                k = j
            }
        }
        // 4.在内层循环结束，也就是找到本轮循环的最小的数以后，再进行交换
        if i != k {
            array[i], array[k] = array[k], array[i]
        }
    }
    return array
}
