package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 399, 0, 34, 4, 3, 2, 0, 1}
	// for i := 0; i < 10; i++ {
	// 	if i >= 0 {
	// 		fmt.Printf("iiiiiiiiiii=%v\n", i)
	// 	}
	// }
	QuickSort(arr, 0, len(arr)-1)
	fmt.Printf("www=%v\n", arr)
}

// 选择排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		temp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp

	}
}

//插入排序
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

// 希尔排序
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
	fmt.Printf("arrarrarrarr=%v\n", arr)
	arrL := len(arr)
	if arrL < 2 {
		return arr
	}
	mid := arrL / 2

	leftArr := MergeSort(arr[:mid])
	fmt.Printf("leftArrleftArrleftArr=%v\n", leftArr)
	rightArr := MergeSort(arr[mid:])
	fmt.Printf("rightArrrightArrrightArr=%v\n", rightArr)
	return merge(leftArr, rightArr)
}
func merge(leftArr, rightArr []int) []int {
	fmt.Printf("mergemergemergemergeleftArrleftArrleftArr=%v\n", leftArr)
	fmt.Printf("mergemergemergemergerightArrrightArrrightArr=%v\n", rightArr)
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
	fmt.Printf("qqqqqqqqqqqq=%v\n", i)
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
