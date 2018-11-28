package main

/*
  快速排序
  递归方式实现
*/
func QuickSort(array []int, start, end int) []int {
    if start < end {
        i, j := start, end
        // 取得关键值
        key := array[(start+end)/2]
        // 循环条件，左边下标大于等于右边下标
        for i <= j {
            // 从前往后比
            // 如果没有比关键值小的，比较下一个
            for array[i] < key {
                i++
            }
            // 从后往前比
            // 如果没有比关键值大的，比较下一个
            for array[j] > key {
                j--
            }
            // 判断下标，调换数据位置
            if i <= j {
                array[i], array[j] = array[j], array[i]
                i++
                j--
            }
        }
        // 此时第一次循环比较结束，关键值的位置已经确定了。左边的值都比关键值小，右边的值都比关键值大，但是两边的顺序还有可能是不一样的，进行下面的递归调用
        if start < j {
            QuickSort(array, start, j)
        }
        if end > i {
            QuickSort(array, i, end)
        }
    }

    return array
}
