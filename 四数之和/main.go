package main

import "sort"

func main() {

}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	result := [][]int{}
	for first := 0; first < numsLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < numsLen; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			aim := target - nums[first] - nums[second]
			four := numsLen - 1
			for three := second + 1; three < numsLen; three++ {
				if three > second+1 && nums[three] == nums[three-1] {
					continue
				}
				for three < four && nums[three]+nums[four]-aim > 0 {
					four--
				}
				if three == four {
					break
				}
				if nums[three]+nums[four] == aim {
					result = append(result, []int{nums[first], nums[second], nums[three], nums[four]})
				}
			}
		}
	}
	return result
}
