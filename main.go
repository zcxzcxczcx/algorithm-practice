package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	format := "20060102"
	// 忽略时分秒
	tStr := time.Now().Format(format)
	fmt.Printf("ffffffffffffffffff=%v\n", tStr)
	tDate, _ := time.Parse(format, tStr)
	fmt.Printf("3333333333333=%v\n", tDate)

}

var text1Map map[string]bool
var max int

func longestCommonSubsequence(text1 string, text2 string) int {
	text1Map = map[string]bool{}
	dfs1([]byte(text1), "")
	fmt.Printf("text1Maptext1Maptext1Map=%v\n", text1Map)
	max = 0
	dfs2([]byte(text2), "")
	return max
}
func dfs1(arr []byte, s string) {

	fmt.Printf("arr=%v\n", arr)
	if s != "" {
		if !text1Map[s] {
			text1Map[s] = true
		}
	}
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		arrCopy := make([]byte, len(arr))
		copy(arrCopy, arr)
		arr1 := arr[i+1:]
		dfs1(arr1, s+string(arrCopy[i]))
		arr = arrCopy
	}
}
func dfs2(arr []byte, s string) {
	if text1Map[s] {
		if len(s) > max {
			max = len(s)
		}
	}
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		arrCopy := make([]byte, len(arr))
		copy(arrCopy, arr)
		arr1 := arr[i+1:]
		dfs2(arr1, s+string(arrCopy[i]))
		arr = arrCopy
	}
}

func jump44(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	result := 0
	steap := 0
	for steap < len(nums)-1 {
		fmt.Printf("dddddd1=%v\n", steap)
		fmt.Printf("nums[steap]nums[steap]=%v\n", nums[steap])
		mx := -1
		for i := 1; i <= nums[steap]; i++ {
			fmt.Printf("iiiiiiiiii=%v\n", i)
			if steap+i+nums[i] > mx {
				mx = steap + i + nums[i]
				steap = steap + i
			}
		}
		fmt.Printf("dddddd2=%v\n", steap)
		result++
	}
	return result
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	i := len(nums) - 1
	for j := i - 1; j >= 0; j-- {
		if nums[j] < i-j {
			continue
		}
		flag := canJump(nums[:j+1])
		if flag {
			return flag
		}
	}

	return false
}

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	result := 0
	flag := false
	var helper func(remainAmount, level, num int)
	min := math.MaxInt64
	helper = func(remainAmount, level, num int) {
		if remainAmount == 0 {
			flag = true
			result = num
			if result < min {
				min = result
			}
			return
		}
		if level < 0 || remainAmount < 0 {
			return
		}
		for i := remainAmount / coins[level]; i >= 0; i-- {
			fmt.Printf("level=%v,ddd=%v\n", level, i)
			copyNum := num
			num = num + i
			helper(remainAmount-i*coins[level], level-1, num)
			num = copyNum

		}
	}
	helper(amount, len(coins)-1, 0)
	if flag {
		return min
	}

	return -1
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 0
	stack := []int{0}
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			for j := 1; j <= nums[stack[i]]; j++ {
				if stack[i]+j > len(nums)-1 {
					break
				}
				if stack[i]+j < len(nums)-1 {
					stack = append(stack, stack[i]+j)
				}
				if stack[i]+j == len(nums)-1 {
					result++
					return result
				}
			}
		}
		stack = stack[l:]
		if len(stack) == 0 {
			return result
		}
		result++

	}
	return result
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
func generateParenthesis(n int) []string {
	var result []string
	strMap := map[string]bool{}
	var helper func(str string, l, r int)

	helper = func(str string, l, r int) {
		if l < 0 || r < 0 {
			return
		}
		if l == 0 && r == 0 {
			if strMap[str] {
				return
			}
			result = append(result, str)
			strMap[str] = true
			return
		}
		if l > r {
			return
		}
		fmt.Printf("dddddd=%v,=%v\n", l, r)
		helper(str+"(", l-1, r)
		helper(str+")", l, r-1)
		return
	}
	helper("", n, n)
	return result
}
