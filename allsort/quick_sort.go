package allsort

func QuickSort(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)

}
func quickSort(arr *[]int, start, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		quickSort(arr, start, pivot-1)
		quickSort(arr, pivot+1, end)
	}

}
func partition(arr *[]int, start, end int) int {
	p := (*arr)[end]
	i := start
	for j := start; j < end; j++ {
		if (*arr)[j] < p {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, i, end)
	return i

}
func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}
