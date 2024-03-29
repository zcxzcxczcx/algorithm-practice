package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 2)
}
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	fmt.Printf("dddddd=%v\n", a)
	fmt.Printf("dddddd=%v\n", b)
	for a != 0 {
		a, b = b%a, a
		fmt.Printf("222222222222=%v\n", a)
		fmt.Printf("4444444444444444444=%v\n", b)
	}
	fmt.Printf("dddddd=%v\n", b)
	return b
}
