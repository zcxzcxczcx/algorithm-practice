package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 1, 5, 2, 6, 9, 2, 5}
	QuickSort(&arr, 0, 7)
	fmt.Printf("dddddddd=%v\n", arr)
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

// 快速排序 用到了分治和递归
func QuickSort(arr *[]int, start, end int) {
	if end <= start {
		return
	}
	cur := partition(arr, start, end)
	QuickSort(arr, start, cur-1)
	QuickSort(arr, cur+1, end)

}
func partition(arr *[]int, start, end int) int {
	cur := start
	for i := start; i <= end; i++ {
		if (*arr)[i] < (*arr)[end] {
			temp := (*arr)[cur]
			(*arr)[cur] = (*arr)[i]
			(*arr)[i] = temp
			cur++
		}
	}
	temp := (*arr)[cur]
	(*arr)[cur] = (*arr)[end]
	(*arr)[end] = temp
	return cur
}
