package main

func BubbleSort(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        for j := 0; j < len(array)-1-i; j++ {
            // 比较数据
            if array[j] < array[j+1] {
                // 数据位置调换
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
    return array
}
