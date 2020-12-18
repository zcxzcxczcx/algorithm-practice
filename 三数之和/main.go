package main

import "sort"

func main() {

}
func threeSum(nums []int) [][]int {
	var ans [][]int
	arrLen := len(nums)
	sort.Ints(nums)

	for first := 0; first < arrLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := arrLen - 1
		for second := first + 1; second < arrLen; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			if second == third {
				break
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}
	return ans
}
