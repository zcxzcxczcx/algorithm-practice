package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 1, 5, 2, 6, 9, 2, 5}
	fmt.Printf("dddddddd=%v\n", MergeSort(arr, 0, 7))
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

// 归并排序
func MergeSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	mid := (left + right) / 2
	arr = MergeSort(arr, left, mid)
	fmt.Printf("dddddddddd111111111111=%v\n", arr)
	arr = MergeSort(arr, mid+1, right)
	fmt.Printf("dddddddddd2222222222=%v\n", arr)
	return merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) []int {
	start := left
	midStart := mid + 1
	l := right - left + 1
	temp := []int{}
	for i := 0; i < l; i++ {
		temp = append(temp, 0)
	}
	k := 0
	for start <= mid && midStart <= right {
		if arr[start] < arr[midStart] {
			temp[k] = arr[start]
			start++
		} else {
			temp[k] = arr[midStart]
			midStart++
		}
		k++
	}
	for start <= mid {
		temp[k] = arr[start]
		start++
		k++
	}
	for midStart <= right {
		temp[k] = arr[midStart]
		midStart++
		k++
	}
	for i := 0; i < l; i++ {
		arr[left+i] = temp[i]
	}

	return arr

}
