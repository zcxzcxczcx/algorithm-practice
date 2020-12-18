package main

func main() {

}
func QuickSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		quickSort(arr, i, len(*arr)-1)
	}

}
func quickSort(arr *[]int, low int, high int) {
	base := (*arr)[low]
	cur := (*arr)[high]
	for low != high {
		if cur < base {
			(*arr)[low] = cur
			low++
			cur = (*arr)[low]
		} else {
			(*arr)[high] = cur
			high--
			cur = (*arr)[high]
		}
	}
	(*arr)[high] = base
}
