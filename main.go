package main

import (
	"algorithm-exercise/allsort"
	"fmt"
)

func main() {
	arr := []int{30, 2, 40, 5, 4, 6, 7, 9, 2, 1}
	allsort.BubbleSort(arr)
	fmt.Printf("ddddddd=%v\n", arr)
}
func numIslands(grid [][]byte) int {
	n := 0
	for k, v := range grid {
		for kv, vv := range v {
			if vv == '1' {
				n++
				dfs(grid, k, kv)
			}
		}
	}
	return n
}
func dfs(grid [][]byte, l, w int) {
	grid[l][w] = '0'
	ll := len(grid)
	wl := 0
	if ll > 0 {
		wl = len(grid[0])
	}

	if l-1 >= 0 && grid[l-1][w] == '1' {
		dfs(grid, l-1, w)
	}
	if w-1 >= 0 && grid[l][w-1] == '1' {
		dfs(grid, l, w-1)
	}
	if w+1 < wl && grid[l][w+1] == '1' {
		dfs(grid, l, w+1)
	}
	if l+1 < ll && grid[l+1][w] == '1' {
		dfs(grid, l+1, w)
	}

}
