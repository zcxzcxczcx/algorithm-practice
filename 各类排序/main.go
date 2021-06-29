package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 1, 5, 200000, 6, 9, 2, 3902, 5}
	radixSort(arr, 6)
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

// 归并排序
func MergeSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	mid := (left + right) / 2
	arr = MergeSort(arr, left, mid)
	arr = MergeSort(arr, mid+1, right)
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

//堆排序
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

// 构建堆
func buildBinaryTree(binaryTree []int) {
	startIndex := (len(binaryTree) - 1) / 2
	for i := startIndex; i >= 0; i-- {
		downAdjust(binaryTree, i)
	}
}

// 上浮节点,添加节点的时候上浮
func upAdjust(binaryTreeArr []int) {
	childIndex := len(binaryTreeArr) - 1
	temp := binaryTreeArr[childIndex]
	parentIndex := (childIndex - 1) / 2
	for temp < binaryTreeArr[parentIndex] {
		binaryTreeArr[childIndex] = binaryTreeArr[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	binaryTreeArr[childIndex] = temp
}

// 计数排序
func countSorting(arr []int, maxValue int) []int {
	arrSort := []int{}
	for i := 0; i < maxValue; i++ {
		arrSort = append(arrSort, 0)
	}
	for i := 0; i < len(arr); i++ {
		arrSort[arr[i]-1]++
	}
	arrSortReturn := []int{}
	for i := 0; i < len(arrSort); i++ {
		if arrSort[i] > 0 {
			for j := 0; j < arrSort[i]; j++ {
				arrSortReturn = append(arrSortReturn, i+1)
			}
		}
	}
	return arrSortReturn
}

// 桶排序
func bucketSort(arr []int, bucketSize int) []int {
	if len(arr) == 0 {
		return nil
	}
	min := arr[0]
	max := arr[0]
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	bucketAccount := ((max - min) / bucketSize) + 1 // 桶的数量
	bucket := [][]int{}
	for i := 0; i < bucketAccount; i++ {
		bucket = append(bucket, []int{})
	}
	for i := 0; i < len(arr); i++ {
		bucketIdx := (arr[i] - min) / bucketSize
		bucket[bucketIdx] = append(bucket[bucketIdx], arr[i])
	}
	arrSort := []int{}
	for i := 0; i < bucketAccount; i++ {
		ar := InsertSort(bucket[i]) // 对每个桶对元素进行排序
		for j := 0; j < len(ar); j++ {
			arrSort = append(arrSort, ar[j])
		}
	}
	return arrSort
}

//基数排序
func radixSort(arr []int, maxDigit int) {
	if len(arr) <= 0 {
		return
	}
	mod := 10
	dev := 1

	for i := 0; i < maxDigit; i++ {
		buket := [10][]int{}
		for j := 0; j < len(arr); j++ {
			buketIndex := (arr[j] % mod) / dev
			buket[buketIndex] = append(buket[buketIndex], arr[j])
		}
		pos := 0
		for j := 0; j < 10; j++ {
			for z := 0; z < len(buket[j]); z++ {
				arr[pos] = buket[j][z]
				pos++
			}
		}
		dev = 10 * dev
		mod = 10 * mod
	}

}
