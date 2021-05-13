package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 1, 5, 2, 6, 9, 2, 5}
	fmt.Printf("dddddddd=%v\n", BubbleSort(arr))
}

// 选择排序
func SelectSort(arr []int) []int {
	arrL := len(arr)
	for i := 0; i < arrL; i++ {
		min := math.MaxInt64
		minK := i
		for k, arrV := range arr[i:] {
			if arrV < min {
				min = arrV
				minK = i + k
			}
		}
		temp := arr[i]
		arr[i] = arr[minK]
		arr[minK] = temp
	}
	return arr
}

// 插入排序
func InsertSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		temp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if temp < arr[j] {
				arr[j+1] = arr[j]
				arr[j] = temp
			} else {
				break
			}
		}
	}
	return arr

}

// 冒泡排序
func BubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}

		}
	}
	return arr
}
