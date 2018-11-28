package main

/*
  插入排序
 */
func InsertSort(array []int) []int {

    for i := 1; i < len(array); i ++ {
        for j := i; j > 0; j-- {
            if array[j-1] > array[j] {
                array[j-1], array[j] = array[j], array[j-1]
            }
        }
    }
    return array

}
