package main

import "math"

func main() {

}
func maxSlidingWindow(nums []int, k int) []int {

	result := []int{}
	numsLen := len(nums)
	if numsLen <= 0 {
		return result
	}
	for i := 0; i < numsLen-k+1; i++ {
		max := math.MinInt64
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
	}
	return result
}
