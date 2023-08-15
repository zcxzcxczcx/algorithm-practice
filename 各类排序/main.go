package main

import (
	"fmt"
)

func main() {
	fmt.Printf("www=%v\n", 1)
	arr := []int{0, 9, 0, 20, 4, 3, 2, 34, 399, 1, 4, 0, 8, 2, 1}
	// for i := 0; i < 10; i++ {
	// 	if i >= 0 {
	// 		fmt.Printf("iiiiiiiiiii=%v\n", i)
	// 	}
	// }
	SelectSort1(arr)
	fmt.Printf("www=%v\n", arr)
}

// 选择排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp
	}
}

// 插入排序
func InsertSort(arr []int) []int {
	arrL := len(arr)
	for i := 0; i < arrL; i++ {
		idx := i
		v := arr[idx]
		for idx > 0 {
			if v > arr[idx-1] {
				break
			}
			arr[idx] = arr[idx-1]
			idx--
		}
		arr[idx] = v
	}
	return arr
}

func InsertSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		idx := i
		for j := i - 1; j >= 0; j-- {
			if temp >= arr[j] {
				idx = j + 1
				break
			}
			arr[j+1] = arr[j]
		}

		arr[idx] = temp
	}
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

// 构建堆
func buildBinaryTree(binaryTree []int) {
	startIndex := (len(binaryTree) - 1) / 2
	for i := startIndex; i >= 0; i-- {
		downAdjust(binaryTree, i)
	}
}

// 堆排序
// 下沉节点:删除的时候下沉
func downAdjust(binaryTreeArr []int, parentIndex int) {
	temp := binaryTreeArr[parentIndex]
	childIndex := 2*parentIndex + 1
	lenth := len(binaryTreeArr)
	for childIndex < lenth {
		if childIndex+1 < lenth && binaryTreeArr[childIndex+1] < binaryTreeArr[childIndex] {
			childIndex = childIndex + 1
		}
		if temp < binaryTreeArr[childIndex] {
			break
		}
		binaryTreeArr[parentIndex] = binaryTreeArr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	binaryTreeArr[parentIndex] = temp
}

// 上浮节点,添加节点的时候上浮
func upAdjust(arr []int) {
	arrL := len(arr)
	childIndex := arrL - 1
	temp := arr[childIndex]
	parentIndex := (childIndex - 1) / 2
	for parentIndex >= 0 {
		if temp >= arr[parentIndex] {
			break
		}
		arr[childIndex] = arr[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	arr[childIndex] = temp
}

// 希尔排序
// 希尔排序又叫缩小增量排序
// 简单插入排序的改进版
func ShellSort(arr []int) {
	arrL := len(arr)
	for gap := arrL / 2; gap >= 1; gap = gap / 2 {
		for i := gap; i < arrL; i++ {
			v := arr[i]
			idx := i
			for idx-gap >= 0 {
				if arr[idx-gap] <= v {
					break
				}
				arr[idx] = arr[idx-gap]
				idx = idx - gap
			}
			arr[idx] = v
		}
	}
}

// 归并排序
func MergeSort(arr []int) []int {
	arrL := len(arr)
	if arrL < 2 {
		return arr
	}
	mid := arrL / 2

	leftArr := MergeSort(arr[:mid])
	rightArr := MergeSort(arr[mid:])
	return merge(leftArr, rightArr)
}
func merge(leftArr, rightArr []int) []int {
	leftL := len(leftArr)
	rightL := len(rightArr)
	tempArr := make([]int, 0)
	leftLIDX := 0
	rightIDX := 0
	for leftLIDX < leftL && rightIDX < rightL {
		if leftArr[leftLIDX] <= rightArr[rightIDX] {
			tempArr = append(tempArr, leftArr[leftLIDX])
			leftLIDX++
		} else {
			tempArr = append(tempArr, rightArr[rightIDX])
			rightIDX++
		}
	}
	if leftLIDX < leftL {
		tempArr = append(tempArr, leftArr[leftLIDX:]...)
	}
	if rightIDX < rightL {
		tempArr = append(tempArr, rightArr[rightIDX:]...)
	}
	return tempArr

}

// 快速排序1 用到了分治和递归
func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	QuickSort(arr, start, i)
	QuickSort(arr, i+1, end)
	return
}
func partition(arr []int, start, end int) int {
	pivotValue := arr[start]
	idx := start + 1
	for i := idx; i <= end; i++ {
		if arr[i] < pivotValue {
			swap(arr, i, idx)
			idx++
		}
	}
	swap(arr, idx-1, start)
	return idx - 1
}
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// 快速排序2 用到了分治和递归
func QuickSort2(arr []int, start, end int) {
	if end <= start {
		return
	}
	cur := partition2(arr, start, end)
	QuickSort2(arr, start, cur-1)
	QuickSort2(arr, cur+1, end)
}
func partition2(arr []int, start, end int) int {
	cur := start
	for i := start; i < end; i++ {
		if arr[i] < arr[end] {
			temp := arr[cur]
			arr[cur] = arr[i]
			arr[i] = temp
			cur++
		}
	}
	temp := arr[cur]
	arr[cur] = arr[end]
	arr[end] = temp
	return cur
}

// 计数排序
func countSorting(arr []int, maxValue int) []int {
	indexArr := make([]int, maxValue+1)
	arrL := len(arr)
	for i := 0; i < arrL; i++ {
		indexArr[arr[i]]++
	}
	resultArr := []int{}
	for k, v := range indexArr {
		for i := 0; i < v; i++ {
			resultArr = append(resultArr, k)
		}
	}
	return resultArr
}

// 桶排序
func bucketSort(arr []int, size int) []int {
	arrL := len(arr)
	if arrL == 0 {
		return nil
	}
	min := arr[0]
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	count := (max-min)/size + 1
	buket := make([][]int, count)

	for i := 0; i < arrL; i++ {
		buket[(arr[i]-min)/size] = append(buket[(arr[i]-min)/size], arr[i])
	}
	result := []int{}
	for i := 0; i < count; i++ {
		ar := MergeSort(buket[i])
		result = append(result, ar...)
	}
	return result
}

// 基数排序
func radixSort(arr []int, maxDigit int) {
	mod := 10
	dev := 1
	for i := 0; i < maxDigit; i++ {
		bucket := make([][]int, 10)
		for j := 0; j < len(arr); j++ {
			bucket[(arr[j]%mod)/dev] = append(bucket[(arr[j]%mod)/dev], arr[j])
		}
		pos := 0
		for j := 0; j < 10; j++ {
			for z := 0; z < len(bucket[j]); z++ {
				arr[pos] = bucket[j][z]
				pos++
			}
		}
		mod = mod * 10
		dev = dev * 10
	}
}
